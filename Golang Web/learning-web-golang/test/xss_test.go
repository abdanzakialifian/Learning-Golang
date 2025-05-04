package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateAutoEscape(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	tempateAutoEscape(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(templateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func tempateAutoEscape(writer http.ResponseWriter, _ *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]any{
		"Title": "Template Auto Escape",
		"Body":  "<p>This is Body</p>",
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateAutoEscapeDisabled(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateAutoEscapeDisabled(writer http.ResponseWriter, _ *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]any{
		"Title": "Template Disabled Escape",
		"Body":  template.HTML("<p>This is Body</p>"),
	})
}

func TestTemplateXSS(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=<p>alert</p>", nil)

	templateXSS(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]any{
		"Title": "Template XSS",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}
