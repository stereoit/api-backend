// +build integration

package mongo

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stereoit/eventival/internal/rand"
	"github.com/stereoit/eventival/pkg/user/domain/repository"

	"github.com/google/uuid"

	"github.com/stereoit/eventival/pkg/user/domain/model"

	"github.com/stretchr/testify/assert"
)

const MONGO_CONNECTION_URI = "mongodb://localhost:27017"

var _repo repository.UserRepository

func getRepository() repository.UserRepository {
	if _repo == nil {
		uri := getenv("MONGODB_URI", MONGO_CONNECTION_URI)
		_repo = createRepository(uri)
	}
	return _repo
}

func createRepository(connectionURI string) repository.UserRepository {
	var testDatabase = "eventival-test-" + rand.String(10)

	opts := &UserRepositoryOpts{
		ConnectionURI: connectionURI,
		Database:      testDatabase,
		Collection:    "users",
	}
	repo, err := NewUserRepository(opts)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// TestMain should make sure the DB is empty
// this is temporary solution as we should mock the
// mongo db response ideally
func TestMain(m *testing.M) {
	connectionURI := flag.String("mongo", MONGO_CONNECTION_URI, "mongo connection string")
	flag.Parse()
	fmt.Printf("connectionURI=%s\n", *connectionURI)
	fmt.Println("Running before all tests")
	_repo = createRepository(*connectionURI)
	code := m.Run()
	// fmt.Println("Running after all tests")
	os.Exit(code)
}

func TestUserRepository_New(t *testing.T) {
	assert := assert.New(t)

	repository := getRepository()

	assert.NotNil(repository)
}

func TestUserRepository_FindAll(t *testing.T) {
	assert := assert.New(t)
	r := getRepository()

	users, err := r.FindAll()
	assert.Nil(err, "Error obtaining all users = %v", err)

	if len(users) != 0 {
		t.Errorf("Users are not empty")
	}

	uid, _ := uuid.NewUUID()
	r.Save(model.NewUser(uid.String(), fmt.Sprintf("%s@example.com", uid.String())))

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
	r := getRepository()

	uid, _ := uuid.NewUUID()
	testUser := model.NewUser(uid.String(), fmt.Sprintf("%s@example.com", uid.String()))
	r.Save(testUser)

	got, err := r.FindByEmail(testUser.GetEmail())
	if err != nil {
		t.Errorf("UserRepository.FindByEmail() error =%v", err)
	}
	assert.Equal(got, testUser)

	got, _ = r.FindByEmail("no-such-email")
	assert.Nil(got, "FindByEmail() should return nil for non existing email")
}

func TestUserRepository_FindByID(t *testing.T) {
	assert := assert.New(t)
	r := getRepository()
	testUser := model.NewUser("1", "email")
	r.Save(testUser)

	got, err := r.FindByID("1")
	if err != nil {
		t.Errorf("UserRepository.FindByID() error =%v", err)
	}
	assert.Equal(got, testUser)

	got, _ = r.FindByID("no-such-id")
	assert.Nil(got, "FindByID() should return nil for non existion ID")
}

func TestUserRepository_Save(t *testing.T) {
	r := getRepository()

	uid, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("cannot create NewUUID() error =%v", err)
	}
	if err = r.Save(model.NewUser(uid.String(), "user@example.com")); err != nil {
		t.Errorf("UserRepository.Save() error =%v", err)
	}
}
