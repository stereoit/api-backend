package registry

import (
	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/pkg/user/domain/repository"
	"github.com/stereoit/eventival/pkg/user/domain/service"
	"github.com/stereoit/eventival/pkg/user/usecase"
)

// Container struct
type Container struct {
	ctn di.Container
}

// NewContainer returns new instance
func NewContainer(userRepository repository.UserRepository) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name: "user-repository",
			Build: func(ctn di.Container) (interface{}, error) {
				return userRepository, nil
			},
		},
		{
			Name:  "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve public method
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean implementation
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("user-repository").(repository.UserRepository)
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
