package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPk(t *testing.T) {
	ctx := context.TODO()
	pk, err := GetPublicKey(ctx)
	if assert.Nil(t, err) {
		assert.NotEmpty(t, pk)
	}
}
