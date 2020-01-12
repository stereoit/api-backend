package registry

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUserRepository(t *testing.T) {
	assert := assert.New(t)
	storageBackendENV := "STORAGE_BACKEND"
	container, _ := NewContainer()

	// test for memory repo
	os.Setenv(storageBackendENV, "MEMORY")
	got := container.ctn.Get("user-repository")
	assert.NotNil(got, "should get memory user repository")

	// test for mongo repo
	os.Setenv(storageBackendENV, "MONGO")
	got = container.ctn.Get("user-repository")
	assert.NotNil(got, "should get mongo user repository")

	// test for default repo
	os.Setenv(storageBackendENV, "")
	got = container.ctn.Get("user-repository")
	assert.NotNil(got, "should get default user repository")

}
