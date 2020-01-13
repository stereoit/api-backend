package v1

import (
	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0/protocol"
	"github.com/stereoit/eventival/pkg/user/usecase"

	"google.golang.org/grpc"
)

// Apply comment
func Apply(server *grpc.Server, ctn di.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Get("user-usecase").(usecase.UserUsecase)))
}
