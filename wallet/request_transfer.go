package wallet

import (
	"context"
	"encoding/json"

	"github.com/fox-one/foxgo/account"
	"github.com/fox-one/foxgo/request"
)

func Transfer(ctx context.Context, token string, pin account.Pin, tr TransferRequest) (*Snapshot, error) {
	data, err := request.Post(ctx, "wallet/transfer",
		account.WithToken(token), account.WithPin(pin), func(p request.Param) error {
			withTransfer(p, tr)
			return nil
		})

	if err != nil {
		return nil, err
	}

	resp := struct {
		Data struct {
			Snapshot Snapshot `json:"snapshot"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(data, &resp)
	return &resp.Data.Snapshot, nil
}
