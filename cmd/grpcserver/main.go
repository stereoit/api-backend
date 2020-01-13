package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/internal/config/services"
	userRPC "github.com/stereoit/eventival/pkg/user/interface/rpc"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create the app container.
	// Do not forget to delete it at the end.
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}

	// load all DI definitions
	err = builder.Add(services.New()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := builder.Build()
	defer app.Delete()

	server := grpc.NewServer()

	userRPC.Apply(server, app)

	go func() {
		log.Printf("start grpc server port: %s", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	server.GracefulStop()
}
