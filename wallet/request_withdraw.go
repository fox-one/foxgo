package wallet

import (
	"context"
	"encoding/json"

	"github.com/fox-one/foxgo/account"
	"github.com/fox-one/foxgo/request"
)

func Withdraw(ctx context.Context, token string, pin account.Pin, wr WithdrawRequest) (*Snapshot, error) {
	data, err := request.Post(ctx, "wallet/withdraw",
		account.WithToken(token), account.WithPin(pin), func(p request.Param) error {
			withWithdraw(p, wr)
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
