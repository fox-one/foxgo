package wallet

import (
	"context"
	"encoding/json"

	"github.com/fox-one/foxgo/account"
	"github.com/fox-one/foxgo/request"
)

func GetAssets(ctx context.Context, token string, includeMainChain bool) (Assets, error) {
	mainchain := 0
	if includeMainChain {
		mainchain = 1
	}

	data, err := request.Get(ctx, "wallet/assets", account.WithToken(token), request.V("entirechain", mainchain))
	if err != nil {
		return nil, err
	}

	resp := struct {
		Data struct {
			Assets Assets `json:"assets"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(data, &resp)
	return resp.Data.Assets, err
}

func GetAssetDetail(ctx context.Context, token, assetid string) (*Asset, error) {
	data, err := request.Get(ctx, "wallet/assets/"+assetid, account.WithToken(token))
	if err != nil {
		return nil, err
	}

	resp := struct {
		Data struct {
			Asset Asset `json:"asset"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(data, &resp)
	return &resp.Data.Asset, err
}
