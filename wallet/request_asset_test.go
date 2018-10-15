package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	token := "dSw1QVosZCxCYyxlLDFnV3h0Wg==.UNW+Tm4jLSMX1AAzROw51foDg0z4RDo8q0nypSr8Bq8="
	if assets, err := GetAssets(context.TODO(), token, true); assert.Nil(t, err) {
		assert.NotEmpty(t, assets)
	}
}

func TestAssetDetail(t *testing.T) {
	token := "dSw1QVosZCxCYyxlLDFnV3h0Wg==.UNW+Tm4jLSMX1AAzROw51foDg0z4RDo8q0nypSr8Bq8="
	if asset, err := GetAssetDetail(context.TODO(), token, ETH); assert.Nil(t, err) {
		assert.NotEmpty(t, asset.PublicKey)
	}
}
