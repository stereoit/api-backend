package registry

import (
	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/pkg/user/domain/service"
	"github.com/stereoit/eventival/pkg/user/interface/persistence/memory"
	"github.com/stereoit/eventival/pkg/user/usecase"
)

// Container struct
type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
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

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo := memory.NewUserRepository()
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
