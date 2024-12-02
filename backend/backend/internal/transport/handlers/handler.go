package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	UserService UserService
	keyJWT      string
	validator   Validator
}

func NewHandler(UserService UserService, keyJWT string, validator Validator) *Handler {
	return &Handler{
		UserService: UserService,
		keyJWT:      keyJWT,
		validator:   validator,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/signup", h.SignUP)
	r.Post("/signin", h.SignIN)

	//	r.With(middlewarejwt.JWTMiddleware(h.keyJWT)).Get("/profile", h.Profile)
	return r
}
