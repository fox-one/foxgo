package wallet

import (
	"context"
	"encoding/json"

	"github.com/fox-one/foxgo/account"
	"github.com/fox-one/foxgo/request"
)

type Status struct {
	IsPinSet bool   `json:"isPinSet"`
	MixinId  string `json:"mixinId"`
	Fullname string `json:"fullname"`
}

func GetWalletStatus(ctx context.Context, token string) (*Status, error) {
	data, err := request.Get(ctx, "wallet/mixinUserInfo", account.WithToken(token))
	if err != nil {
		return nil, err
	}

	resp := struct {
		Data struct {
			Status `json:"status"`
		} `json:"data"`
	}{}

	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp.Data.Status, nil
}
