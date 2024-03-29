package registry

import (
	"github.com/stereoit/eventival/internal/env"
	"log"
	"strings"

	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/pkg/user/domain/repository"
	"github.com/stereoit/eventival/pkg/user/domain/service"
	"github.com/stereoit/eventival/pkg/user/interface/persistence/memory"
	"github.com/stereoit/eventival/pkg/user/interface/persistence/mongo"
	"github.com/stereoit/eventival/pkg/user/usecase"
)

// Services for user module
var Services = []di.Def{
	{
		Name:  "user-repository",
		Build: buildUserRepository,
	},
	{
		Name:  "user-usecase",
		Build: buildUserUsecase,
	},
}

func buildUserRepository(ctn di.Container) (interface{}, error) {
	var (
		userRepository repository.UserRepository
		err            error
	)

	// get the storage implementation
	storageType := strings.ToUpper(env.Get("STORAGE_BACKEND", "memory"))
	switch storageType {
	case "MEMORY":
		userRepository = memory.NewUserRepository()
	case "MONGO":
		userRepository, err = mongo.NewUserRepository(&mongo.UserRepositoryOpts{
			ConnectionURI: env.Get("MONGO_URI", "mongodb://localhost:27017"),
			Database:      env.Get("USER_DB", "eventival"),
			Collection:    env.Get("USER_COLLECTION", "users"),
		})
		if err != nil {
			log.Fatalf("Failed to initialize storage: %v\n", err)
			return nil, err
		}
	default:
		userRepository = memory.NewUserRepository()
	}

	return userRepository, nil
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("user-repository").(repository.UserRepository)
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
