package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sarulabs/di"
	"github.com/stereoit/eventival/internal/config/services"
	"github.com/stereoit/eventival/internal/env"
	userRestService "github.com/stereoit/eventival/pkg/user/interface/rest"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	port := env.Get("REST_PORT", "8000")
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

	// build DI container
	err = builder.Add(services.New()...)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := builder.Build()
	defer app.Delete()

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	userRestService.Apply("/users", r, app)

	server := http.Server{
		Handler: r,
	}

	go func() {
		log.Printf("starting rest server port: %s", lis.Addr())
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping rest server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("reset server shutdown: %v", err)
	}
}
