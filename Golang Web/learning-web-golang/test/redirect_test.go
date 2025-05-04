package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", redirectFrom)
	mux.HandleFunc("/redirect-to", redirectTo)
	mux.HandleFunc("/redirect-out", redirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func redirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

func redirectFrom(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func redirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://google.com", http.StatusTemporaryRedirect)
}
