package rpc

import (
	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0"

	"google.golang.org/grpc"
)

// Apply this interface to given server
func Apply(server *grpc.Server, ctn di.Container) {
	v1.Apply(server, ctn)
}
