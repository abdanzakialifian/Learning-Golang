package test

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TestTemplateCaching(t *testing.T) {
	record := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateCaching(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var embedTemplates embed.FS

var myTemplates = template.Must(template.ParseFS(embedTemplates, "templates/*.gohtml"))

func templateCaching(writer http.ResponseWriter, _ *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}
