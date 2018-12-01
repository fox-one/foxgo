package wallet

import (
	"context"
	"encoding/json"

	"github.com/fox-one/foxgo/account"
	"github.com/fox-one/foxgo/request"
)

func QuerySnapshots(ctx context.Context, token, assetId string, desc bool, limit int, cursor string) ([]Snapshot, *request.Pagination, error) {
	order := "ASC"
	if desc {
		order = "DESC"
	}

	data, err := request.Get(ctx, "wallet/snapshots",
		account.WithToken(token),
		request.V("assetId", assetId),
		request.V("order", order),
		request.V("limit", limit),
		request.V("cursor", cursor),
	)

	if err != nil {
		return nil, nil, err
	}

	resp := struct {
		Data struct {
			Snapshots []Snapshot         `json:"snapshots"`
			Cursor    request.Pagination `json:"pagination"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(data, &resp)
	return resp.Data.Snapshots, &resp.Data.Cursor, err
}
