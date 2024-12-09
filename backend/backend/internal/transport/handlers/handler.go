package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler struct {
	AuthService   AuthService
	WalletService WalletService
	UserService   UserService
	keyJWT        string
	Validator     Validator
	jwtMiddleware JWTMiddleware
}

type JWTMiddleware interface {
	JWTMiddlewareCustomer() func(http.Handler) http.Handler
	JWTMiddlewareAdmin() func(http.Handler) http.Handler
}

func NewHandler(UserService UserService, keyJWT string, validator Validator, jwtMiddleware JWTMiddleware, AuthService AuthService, WalletService WalletService) *Handler {
	return &Handler{
		UserService:   UserService,
		WalletService: WalletService,
		AuthService:   AuthService,
		keyJWT:        keyJWT,
		Validator:     validator,
		jwtMiddleware: jwtMiddleware,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	//For all
	r.Post("/signup", h.SignUP)
	r.Post("/signin", h.SignIN)

	//For users
	r.With(h.jwtMiddleware.JWTMiddlewareCustomer()).Get("/wallet/balance", h.WalletBalance)
	r.With(h.jwtMiddleware.JWTMiddlewareCustomer()).Put("/update/{id}", h.UserUpdate)

	//For admins

	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{username}", h.UserFindByUsername)
	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{email}", h.UserFindByEmail)
	return r
}
