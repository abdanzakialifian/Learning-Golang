package middleware

import (
	"learning-database-migration-golang/exception"
	"net/http"
)

type ErrorMiddleware struct {
	Handler http.Handler
}

func NewErrorMiddleware(handler http.Handler) *ErrorMiddleware {
	return &ErrorMiddleware{Handler: handler}
}

func (middleware *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			exception.ErrorHandler(writer, request, err)
		}
	}()
	middleware.Handler.ServeHTTP(writer, request)
}
