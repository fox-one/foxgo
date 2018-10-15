package account

import (
	"context"
	"encoding/json"

	"github.com/fox-one/foxgo/request"
)

type User struct {
	ID          uint   `json:"userId"`
	MixinUserId string `json:"mixinUserId"`
	Fullname    string `json:"fullname"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Language    string `json:"language"`
	IsPinSet    bool   `json:"isPinSet"`
}

func GetUserDetail(ctx context.Context, token string) (*User, error) {
	data, err := request.Get(ctx, "account/detail", WithToken(token))

	if err != nil {
		return nil, err
	}

	resp := struct {
		Data struct {
			User User `json:"user"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(data, &resp)
	return &resp.Data.User, nil
}
