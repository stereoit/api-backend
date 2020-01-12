package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnv(t *testing.T) {
	assert := assert.New(t)

	// test of fallback
	want := "default"
	if got := Get("NONEXISTENT", "default"); got != want {
		t.Errorf("getEnv() = %q, want %q", got, want)
	}

	// test of reading ENV ok
	want = "SUCCESS"
	os.Setenv("DUMMY", want)
	got := Get("DUMMY", "FAIL")
	assert.Equal(want, got, "we should be able to read ENV properly")
	os.Setenv("DUMMY", "")
}
