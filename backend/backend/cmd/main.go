package main

import (
	"context"
	"e-commerce/backend/config"
	"e-commerce/backend/internal/repository"
	"e-commerce/backend/internal/service"
	"e-commerce/backend/internal/transport"
	"e-commerce/backend/internal/transport/handlers"
	"e-commerce/backend/internal/transport/middlewarejwt"
	"e-commerce/backend/internal/util/validator"
	"e-commerce/backend/pkg/db"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func LoadEnv() {
	err := godotenv.Load("D:/e-commerce/backend/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env")
	}
}

func main() {

	LoadEnv()

	keyJWT := os.Getenv("JWT_SECRET_KEY")
	fmt.Println("JWT Key:", keyJWT)
	if keyJWT == "" {
		log.Fatal().Msg("JWT secret key is not configured")
	}

	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	database, err := db.NewGormDB(cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect database")
	}

	walletRepo, err := repository.NewWalletRepository(database)
	if err != nil {
		log.Error().Err(err).Msg("error creating wallet repository")
	}

	authRepo, err := repository.NewAuthRepository(database, walletRepo)
	if err != nil {
		log.Error().Err(err).Msg("error creating auth repository")
	}

	userRepo, err := repository.NewUserRepository(database)
	if err != nil {
		log.Error().Err(err).Msg("error creating user repository")
	}

	authServ := service.NewAuthService(authRepo, walletRepo)
	walletServ := service.NewWalletService(walletRepo)
	userServ := service.NewUserService(userRepo)

	valid := validator.NewGoValidator()

	jwtMiddleware := middlewarejwt.NewJWTMiddleware(keyJWT) //

	h := handlers.NewHandler(userServ, keyJWT, valid, jwtMiddleware, authServ, walletServ) //jwtmiddleware
	r := h.InitRoutes()

	srv := transport.NewServer(cfg, r)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(); err != nil && http.ErrServerClosed == nil {
			log.Error().Err(err).Msg("Server failed to start")
			fmt.Printf("Server failed: %v\n", err)
		}
	}()

	<-stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Error during shutdown: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
