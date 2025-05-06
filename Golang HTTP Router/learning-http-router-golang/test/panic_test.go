package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic : ", error)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Panic : Ups", string(body))
}
