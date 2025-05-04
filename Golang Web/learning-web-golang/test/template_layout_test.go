package test

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateLayout(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateLayout(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templateGoHtml embed.FS

func templateLayout(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFS(templateGoHtml, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "layout", map[string]any{
		"Title": "Template Layout",
		"Name":  "Zaki",
	})
}
