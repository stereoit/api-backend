package rpc

import (
	"github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0"
	"github.com/stereoit/eventival/pkg/user/registry"

	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
