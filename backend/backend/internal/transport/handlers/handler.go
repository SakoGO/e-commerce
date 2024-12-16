package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler struct {
	AuthService     AuthService
	UserService     UserService
	ShopService     ShopService
	CategoryService CategoryService
	ProductService  ProductService
	keyJWT          string
	validator       Validator
	jwtMiddleware   JWTMiddleware
}

type JWTMiddleware interface {
	JWTMiddlewareUser() func(http.Handler) http.Handler
	JWTMiddlewareAdmin() func(http.Handler) http.Handler
}

func NewHandler(UserService UserService, keyJWT string, validator Validator, jwtMiddleware JWTMiddleware,
	AuthService AuthService, ShopService ShopService, ProductService ProductService, CategoryService CategoryService) *Handler {

	return &Handler{
		UserService:     UserService,
		AuthService:     AuthService,
		ShopService:     ShopService,
		ProductService:  ProductService,
		CategoryService: CategoryService,
		keyJWT:          keyJWT,
		validator:       validator,
		jwtMiddleware:   jwtMiddleware,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	//auth
	r.Post("/signup", h.SignUP)
	r.Post("/signin", h.SignIN)

	//shop
	r.Get("/get/shop/{id}", h.GetShopID)
	r.Get("/get/product/{id}", h.GetProductByID)
	r.Get("/get/product/byShopID/{id}", h.GetProductsByShopID)

	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/create_shop", h.CreateShop)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/create_product", h.CreateProduct)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/update/{id}", h.UpdateShop)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/delete/{id}", h.DeleteShop)

	//product
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/product/update/{id}", h.UpdateProduct)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/product/delete/{id}", h.DeleteProduct)

	//For admins
	r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Post("/category/create", h.CreateCategory)
	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{username}", h.UserFindByUsername)
	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{email}", h.UserFindByEmail)
	return r
}
