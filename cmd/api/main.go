package main

import (
	"auth/cmd/api/handlers"
	"auth/internal/sys"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var version = "SNAPSHOT 0.0.1"

func main() {
	log.Println("Starting...")
	log.Printf("VERSION : %s", version)
	ctx := context.Background()
	// Parsing configs from env file
	log.Println("Parsing configs...")
	conf, err := sys.NewConfigWithEnv()
	if err != nil {
		log.Fatalf("Failed to parse config : %v\n", err)
	}

	log.Printf("Starting %s environment\n", conf.Env)

	log.Println("Initializing logger...")
	logger, err := sys.Logger(conf.Env)
	if err != nil {
		log.Fatalf("Failed to create logger : %v", err)
	}

	log.Println("Initializing application...")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	db, err := sys.Postgres(ctx, conf.Postgres.DSN)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	app, err := handlers.API(shutdown, logger, db, conf.Token)
	if err != nil {
		log.Fatalf("Cannot initialize application : %v", err)
	}

	log.Println("Starting server...")
	server := http.Server{
		Addr:    ":" + conf.Port,
		Handler: app,
	}

	serverErrChan := make(chan error, 1)

	go func() {
		serverErrChan <- server.ListenAndServe()
	}()

	log.Printf("Started the server on %s", conf.Port)

	select {
	case <-serverErrChan:
		log.Fatalf("server error: %v", err)
	case <-shutdown:
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Error when server.Shutdown() : %v\n", err)
			server.Close()
			return
		}
		log.Println("Server was gracefully stopped")
	}
}
