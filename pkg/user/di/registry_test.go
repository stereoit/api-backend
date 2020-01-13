package registry

import (
	"log"
	"os"
	"testing"

	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
)

func createContainer() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = builder.Add(Services...)
	if err != nil {
		log.Fatal(err.Error())
	}

	return builder.Build()
}

func Test_GetUserRepositoryMemory(t *testing.T) {
	assert := assert.New(t)
	container := createContainer()

	// test for memory repo
	os.Setenv("STORAGE_BACKEND", "MEMORY")
	got := container.Get("user-repository")
	assert.NotNil(got, "should get memory user repository")
}

func Test_GetUserRepositoryMongo(t *testing.T) {
	assert := assert.New(t)
	container := createContainer()
	os.Setenv("STORAGE_BACKEND", "MONGO")
	got := container.Get("user-repository")
	assert.NotNil(got, "should get mongo user repository")
}

func Test_GetUserRepositoryDefault(t *testing.T) {
	assert := assert.New(t)
	container := createContainer()
	os.Setenv("STORAGE_BACKEND", "")
	got := container.Get("user-repository")
	assert.NotNil(got, "should get default user repository")
}

func Test_UserUsecase(t *testing.T) {
	assert := assert.New(t)
	container := createContainer()

	got := container.Get("user-usecase")
	assert.NotNil(got, "shold get user usecase")
}
