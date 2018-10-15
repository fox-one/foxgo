package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	token      = "dSw1QVosZCxCYyxlLDFnV3h0Wg==.UNW+Tm4jLSMX1AAzROw51foDg0z4RDo8q0nypSr8Bq8="
	wrongToken = "adlsakhfdalkhfjkh"
)

func TestGetUserDetail(t *testing.T) {
	ctx := context.TODO()
	u, err := GetUserDetail(ctx, token)
	if assert.Nil(t, err) {
		assert.NotZero(t, u.ID)
	}
}

func TestLoginRequired(t *testing.T) {
	_, err := GetUserDetail(context.TODO(), wrongToken)
	assert.True(t, IsLoginRequired(err))
}
