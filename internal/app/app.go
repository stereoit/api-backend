package app

import (
	"log"
	"os"
	"strings"

	"github.com/stereoit/eventival/pkg/user/domain/repository"

	"github.com/stereoit/eventival/pkg/user/interface/persistence/memory"
	"github.com/stereoit/eventival/pkg/user/interface/persistence/mongo"
)

// GetUserRepository builder pattern
func GetUserRepository() (repository.UserRepository, error) {

	var (
		userRepository repository.UserRepository
		err            error
	)

	// get the storage implementation
	storageType := strings.ToUpper(GetEnv("STORAGE_BACKEND", "memory"))
	switch storageType {
	case "MEMORY":
		userRepository = memory.NewUserRepository()
	case "MONGO":
		userRepository, err = mongo.NewUserRepository(&mongo.UserRepositoryOpts{
			ConnectionURI: GetEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database:      GetEnv("USER_DB", "eventival"),
			Collection:    GetEnv("USER_COLLECTION", "users"),
		})
		if err != nil {
			log.Fatalf("Failed to initialize storage: %v\n", err)
		}
		log.Println("Storage set to mongo")
	default:
		userRepository = memory.NewUserRepository()
	}

	return userRepository, nil
}

// GetEnv read ENV variable with fallback to default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
