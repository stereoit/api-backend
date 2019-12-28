package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnv(t *testing.T) {
	assert := assert.New(t)

	// test of fallback
	want := "default"
	if got := GetEnv("NONEXISTENT", "default"); got != want {
		t.Errorf("getEnv() = %q, want %q", got, want)
	}

	// test of reading ENV ok
	want = "SUCCESS"
	os.Setenv("DUMMY", want)
	got := GetEnv("DUMMY", "FAIL")
	assert.Equal(want, got, "we should be able to read ENV properly")
	os.Setenv("DUMMY", "")
}

func Test_GetUserRepository(t *testing.T) {
	assert := assert.New(t)
	storageBackendENV := "STORAGE_BACKEND"

	// test for memory repo
	os.Setenv(storageBackendENV, "MEMORY")
	got, _ := GetUserRepository()
	assert.NotNil(got, "should get memory user repository")

	// test for mongo repo
	os.Setenv(storageBackendENV, "MONGO")
	got, _ = GetUserRepository()
	assert.NotNil(got, "should get mongo user repository")

	// test for default repo
	os.Setenv(storageBackendENV, "")
	got, _ = GetUserRepository()
	assert.NotNil(got, "should get default user repository")

}
