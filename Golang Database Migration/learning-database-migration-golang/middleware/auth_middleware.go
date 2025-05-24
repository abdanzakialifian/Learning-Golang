package middleware

import (
	"learning-database-migration-golang/exception"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := request.Header.Get("X-API-Key")

	if apiKey == "secret" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		panic(exception.NewUnauthorized("Token Not Available"))
	}
}
