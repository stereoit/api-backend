package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/stereoit/eventival/internal/app"
	"github.com/stereoit/eventival/pkg/user/interface/rpc"

	"google.golang.org/grpc"

	"github.com/stereoit/eventival/pkg/user/registry"
)

func main() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// get the storage implementation
	userStorage, err := app.GetUserRepository()
	if err != nil {
		log.Fatalf("failed to get user repository: %v", err)
	}
	ctn, err := registry.NewContainer(userStorage)
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	server := grpc.NewServer()

	rpc.Apply(server, ctn)

	go func() {
		log.Printf("start grpc server port: %s", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	server.GracefulStop()
	ctn.Clean()
}
