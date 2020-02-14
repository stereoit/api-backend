package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextKey(t *testing.T) {
	assert := assert.New(t)
	expected := "eventival test"
	assert.Equal(expected, contextKey("test").String(), "context key should match")
}
