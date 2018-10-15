package request

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/fox-one/foxgo/foxerr"
)

var baseApi = "api.fox.one/api"

func ActiveDevEnvironment() {
	baseApi = "dev.fox.one/api"
}

type Request struct {
	method string
	uri    *url.URL

	param Param
}

func New(method, endpoint string, param Param) (*Request, error) {
	u, err := url.Parse(path.Join(baseApi, endpoint))
	if err != nil {
		return nil, err
	}

	u.Scheme = "https"

	return &Request{
		method: method,
		uri:    u,
		param:  param,
	}, nil
}

func (r *Request) request() (*http.Request, error) {
	u := r.uri

	var body io.Reader = nil
	var header http.Header = nil

	if p := r.param; p != nil {
		if query := u.Query(); len(query) == 0 {
			u.RawQuery = p.Query().Encode()
		} else {
			for key, values := range p.Query() {
				for _, v := range values {
					query.Add(key, v)
				}
			}
			u.RawQuery = query.Encode()
		}

		body = bytes.NewBuffer(p.Body())
		header = p.Header()
	}

	req, err := http.NewRequest(r.method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header = header
	return req, err
}

func (r *Request) Do(ctx context.Context) ([]byte, error) {
	req, err := r.request()
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := sharedClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	var foxErr = struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Hint string `json:"hint"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&foxErr); err != nil {
		return nil, err
	}

	if foxErr.Code != 0 {
		foxErr := foxerr.New(foxErr.Code, foxErr.Msg, foxErr.Hint)
		return nil, foxerr.NewRequestFailure(foxErr, resp.StatusCode)
	}

	return nil, errors.New(resp.Status)
}

type BuildParamFunc func(p Param) error

// V return a build param func that set values
func V(key string, value interface{}) BuildParamFunc {
	return func(p Param) error {
		p.SetValue(key, value)
		return nil
	}
}

// H return a build param func that set header
func H(key, value string) BuildParamFunc {
	return func(p Param) error {
		p.SetHeader(key, value)
		return nil
	}
}

func buildQueryParam(funcs ...BuildParamFunc) (Param, error) {
	if len(funcs) == 0 {
		return nil, nil
	}

	param := NewQueryParam()
	for _, f := range funcs {
		if err := f(param); err != nil {
			return nil, err
		}
	}

	return param, nil
}

func buildJsonParam(funcs ...BuildParamFunc) (Param, error) {
	if len(funcs) == 0 {
		return nil, nil
	}

	param := NewJsonParam()
	for _, f := range funcs {
		if err := f(param); err != nil {
			return nil, err
		}
	}

	return param, nil
}

func Get(ctx context.Context, endpoint string, funcs ...BuildParamFunc) ([]byte, error) {
	param, err := buildQueryParam(funcs...)
	if err != nil {
		return nil, err
	}

	r, err := New(http.MethodGet, endpoint, param)
	if err != nil {
		return nil, err
	}

	return r.Do(ctx)
}

func Put(ctx context.Context, endpoint string, funcs ...BuildParamFunc) ([]byte, error) {
	param, err := buildJsonParam(funcs...)
	if err != nil {
		return nil, err
	}

	r, err := New(http.MethodPut, endpoint, param)
	if err != nil {
		return nil, err
	}

	return r.Do(ctx)
}

func Post(ctx context.Context, endpoint string, funcs ...BuildParamFunc) ([]byte, error) {
	param, err := buildJsonParam(funcs...)
	if err != nil {
		return nil, err
	}

	r, err := New(http.MethodPost, endpoint, param)
	if err != nil {
		return nil, err
	}

	return r.Do(ctx)
}

func Delete(ctx context.Context, endpoint string, funcs ...BuildParamFunc) ([]byte, error) {
	param, err := buildQueryParam(funcs...)
	if err != nil {
		return nil, err
	}

	r, err := New(http.MethodDelete, endpoint, param)
	if err != nil {
		return nil, err
	}

	return r.Do(ctx)
}
