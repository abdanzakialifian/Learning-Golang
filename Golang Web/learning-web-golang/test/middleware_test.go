package test

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Middleware")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Middleware")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Error Handler")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Executed Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Error Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Executed")
		panic("Ups Error!")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	errorHandler := ErrorHandler{
		Handler: &logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &errorHandler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
