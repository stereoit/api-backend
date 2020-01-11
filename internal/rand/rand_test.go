package rand

import "testing"

import "github.com/stretchr/testify/assert"

func Test_String(t *testing.T) {
	assert := assert.New(t)
	length := 10

	got := String(length)
	assert.NotNil(got, "generated text should not be nil")
	assert.Len(got, length, "generated string should match given length")
}

func Test_StringWithCharset(t *testing.T) {
	assert := assert.New(t)
	length := 5
	charset := "A"
	expected := "AAAAA"

	got := StringWithCharset(length, charset)
	assert.NotNil(got, "generated text should not be nil")
	assert.Len(got, length, "generated string should match given length")
	assert.Equal(expected, got, "should equal")
}
