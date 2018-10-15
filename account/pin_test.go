package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPin(t *testing.T) {
	ctx := context.TODO()

	pk, err := GetPublicKey(ctx)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	token := "dSw1QVosZCxCYyxlLDFnV3h0Wg==.UNW+Tm4jLSMX1AAzROw51foDg0z4RDo8q0nypSr8Bq8="
	user, err := GetUserDetail(ctx, token)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	if !user.IsPinSet {
		// set pin
		err = UpdatePin(ctx, token, EmptyPin, NewPin("123456", pk))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
	}

	// change pin to 654321
	err = UpdatePin(ctx, token, NewPin("123456", pk), NewPin("654321", pk))
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// change pin back to 123456
	err = UpdatePin(ctx, token, NewPin("654321", pk), NewPin("123456", pk))
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
}
