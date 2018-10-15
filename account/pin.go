package account

import (
	"context"

	"github.com/fox-one/foxgo/request"
)

type Pin struct {
	code string
	pk   string
}

var EmptyPin = Pin{}

func NewPin(code, pk string) Pin {
	return Pin{code: code, pk: pk}
}

func VerifyPin(ctx context.Context, token string, pin Pin) (bool, error) {
	if _, err := request.Put(ctx, "account/pin-verify", WithToken(token), WithPin(pin)); err == nil {
		return true, nil
	} else if IsPinInvalid(err) {
		return false, nil
	} else {
		return false, err
	}
}

// Use EmptyPin as pin when set pin first time
func UpdatePin(ctx context.Context, token string, pin, newPin Pin) error {
	_, err := request.Put(ctx, "account/pin", WithToken(token), WithPin(pin), WithNewPin(newPin))
	return err
}
