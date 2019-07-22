package memory_test

import (
	"testing"

	"github.com/stereoit/eventival/pkg/user/domain/model"
	storage "github.com/stereoit/eventival/pkg/user/interface/persistence/memory"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_New(t *testing.T) {
	assert := assert.New(t)

	repository := storage.NewUserRepository()
	assert.NotNil(repository)
}

func TestUserRepository_FindAll(t *testing.T) {
	r := storage.NewUserRepository()

	users, err := r.FindAll()
	if err != nil {
		t.Fatalf("Error obtaining all users = %v", err)
	}

	if len(users) != 0 {
		t.Errorf("Users are not empty")
	}

	r.Save(model.NewUser("1", "email", "firstName", "lastName"))
	users, err = r.FindAll()
	if err != nil {
		t.Fatalf("Error obtaining all users = %v", err)
	}

	if len(users) != 1 {
		t.Errorf("FindAll() size = %v, want %v", users, 1)
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	assert := assert.New(t)
	r := storage.NewUserRepository()
	testUser := model.NewUser("1", "email", "firstName", "lastName")
	r.Save(testUser)

	got, err := r.FindByEmail("email")
	if err != nil {
		t.Errorf("UserRepository.FindByEmail() error =%v", err)
	}
	assert.Equal(got, testUser)

	got, _ = r.FindByEmail("no-such-email")
	assert.Nil(got, "FindByEmail() should return nil for non existing email")
}

func TestUserRepository_Save(t *testing.T) {
	r := storage.NewUserRepository()
	if err := r.Save(model.NewUser("1", "email", "firstName", "lastName")); err != nil {
		t.Errorf("UserRepository.Save() error =%v", err)
	}
}
