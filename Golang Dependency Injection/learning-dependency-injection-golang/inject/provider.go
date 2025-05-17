package inject

import (
	"learning-dependency-injection-golang/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewServer(middleware *middleware.ErrorMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: middleware,
	}
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewMiddleware(router *httprouter.Router) *middleware.ErrorMiddleware {
	authMiddleware := middleware.NewAuthMiddleware(router)
	errorMiddleware := middleware.NewErrorMiddleware(authMiddleware)
	return errorMiddleware
}
