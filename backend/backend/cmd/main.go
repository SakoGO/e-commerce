package main

import (
	"context"
	"e-commerce/backend/config"
	"e-commerce/backend/internal/transport"
	"fmt"
	"github.com/go-chi/chi/v5"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	r := chi.NewRouter()

	srv := transport.NewServer(cfg, r)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(); err != nil {
			fmt.Printf("Server failed: %v\n", err)
		}
	}()

	<-stop
	fmt.Println("Gracefully shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Error during shutdown: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
