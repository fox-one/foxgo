package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Param interface {
	// read url query
	Query() url.Values
	// read body
	Body() []byte
	// read header
	Header() http.Header
	// Set Value
	SetValue(key string, value interface{})
	// Set Header
	SetHeader(key, value string)
}

// type param map[string]interface{}
type param struct {
	values map[string]interface{}
	header http.Header
}

func (p param) SetValue(key string, value interface{}) {
	p.values[key] = value
}

func (p param) SetHeader(key, value string) {
	p.header.Set(key, value)
}

func (p param) Header() http.Header {
	return p.header
}

type jsonParam struct {
	param
}

func (jp jsonParam) Query() url.Values {
	return nil
}

func (jp jsonParam) Body() []byte {
	b, _ := json.Marshal(jp.param.values)
	return b
}

var _ Param = &jsonParam{}

type queryParam struct {
	param
}

func (qp queryParam) Query() url.Values {
	query := url.Values{}
	for k, v := range qp.param.values {
		query.Set(k, fmt.Sprintf("%v", v))
	}

	return query
}

func (qp queryParam) Body() []byte {
	return nil
}

var _ Param = &queryParam{}

func newParam() param {
	return param{
		values: make(map[string]interface{}),
		header: make(http.Header),
	}
}

func NewJsonParam() Param {
	return &jsonParam{param: newParam()}
}

func NewQueryParam() Param {
	return &queryParam{param: newParam()}
}
