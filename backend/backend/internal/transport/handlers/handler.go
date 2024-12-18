package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	ShopHandler     *ShopHandler
	AuthHandler     *AuthHandler
	ProductHandler  *ProductHandler
	CategoryHandler *CategoryHandler
	keyJWT          string
	validator       Validator
	jwtMiddleware   JWTMiddleware
}

func NewHandler(keyJWT string, validator Validator, jwtMiddleware JWTMiddleware, ShopService ShopService, AuthService AuthService, ProductService ProductService, CategoryService CategoryService) *Handler {

	return &Handler{
		AuthHandler:     NewAuthHandler(AuthService, validator),
		ShopHandler:     NewShopHandler(ShopService, validator),
		ProductHandler:  NewProductHandler(ProductService, validator),
		CategoryHandler: NewCategoryHandler(CategoryService),
		keyJWT:          keyJWT,
		validator:       validator,
		jwtMiddleware:   jwtMiddleware,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	//auth
	r.Post("/signup", h.AuthHandler.SignUP)
	r.Post("/signin", h.AuthHandler.SignIN)

	//shop
	r.Get("/get/shop/{id}", h.ShopHandler.GetShopID)
	r.Get("/get/product/{id}", h.ProductHandler.GetProductByID)
	r.Get("/get/product/byShopID/{id}", h.ProductHandler.GetProductsByShopID)

	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/create_shop", h.ShopHandler.CreateShop)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/shop/update/{id}", h.ShopHandler.UpdateShop)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Delete("/shop/delete/{id}", h.ShopHandler.DeleteShop)

	//product
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/product/create_product", h.ProductHandler.CreateProduct)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Post("/product/update", h.ProductHandler.UpdateProduct)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Delete("/product/delete/{id}", h.ProductHandler.DeleteProduct)
	r.With(h.jwtMiddleware.JWTMiddlewareUser()).Delete("/product/delete/shop/{id}", h.ProductHandler.DeleteProductsByShopID)

	//For admins
	r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Post("/category/create", h.CategoryHandler.CreateCategory)
	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{username}", h.UserFindByUsername)
	// r.With(h.jwtMiddleware.JWTMiddlewareAdmin()).Get("/user/{email}", h.UserFindByEmail)
	return r
}
