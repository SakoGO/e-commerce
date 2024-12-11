package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler struct {
	AuthService   AuthService
	UserService   UserService
	ShopService   ShopService
	keyJWT        string
	validator     Validator
	jwtMiddleware JWTMiddleware
}

type JWTMiddleware interface {
	JWTMiddlewareUser() func(http.Handler) http.Handler
	JWTMiddlewareAdmin() func(http.Handler) http.Handler
}

func NewHandler(UserService UserService, keyJWT string, validator Validator, jwtMiddleware JWTMiddleware,
	AuthService AuthService, ShopService ShopService) *Handler {

	return &Handler{
		UserService:   UserService,
		AuthService:   AuthService,
		ShopService:   ShopService,
		keyJWT:        keyJWT,
		validator:     validator,
		jwtMiddleware: jwtMiddleware,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	//For all
	r.Post("/signup", h.SignUP)
	r.Post("/signin", h.SignIN)

	//For Users
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/create_shop", h.CreateShop)

	//For admins

	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{username}", h.UserFindByUsername)
	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{email}", h.UserFindByEmail)
	return r
}
