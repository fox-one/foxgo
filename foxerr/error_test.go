package foxerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseError(t *testing.T) {
	err := newBaseError(100, "somthing")
	assert.NotNil(t, err)
	assert.Equal(t, 100, err.code)
	assert.NotEmpty(t, err.msg)
	assert.Empty(t, err.hint)
}

func TestNewBaseErrorWithHint(t *testing.T) {
	err := newBaseError(100, "err", "hint", "blablabla")
	assert.NotNil(t, err)
	assert.Equal(t, 100, err.code)
	assert.Equal(t, "err", err.msg)
	assert.Equal(t, "hint", err.hint)
}
