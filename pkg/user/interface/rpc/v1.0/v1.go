package v1

import (
	"github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0/protocol"
	"github.com/stereoit/eventival/pkg/user/registry"
	"github.com/stereoit/eventival/pkg/user/usecase"
	"google.golang.org/grpc"
)

// Apply comment
func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
