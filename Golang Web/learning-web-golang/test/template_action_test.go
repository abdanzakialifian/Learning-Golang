package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateActionIf(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateActionIf(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateActionIf(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/if_template.gohtml"))
	t.ExecuteTemplate(writer, "if_template.gohtml", map[string]any{
		"Title": "Template Action",
		"Name":  "Zaki",
	})
}

func TestTemplateComparator(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateComparator(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateComparator(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]any{
		"Title":      "Template Action",
		"FinalValue": 65,
	})
}

func TestTemplateRange(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateRange(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateRange(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]any{
		"Title": "Template Action",
		"Hobbies": []string{
			"Game",
			"Read",
			"Code",
		},
	})
}

func TestTemplateNested(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateNested(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateNested(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]any{
		"Title": "Template Action",
		"Address": map[string]any{
			"Street": "Jl. Dipayuda Utara No.23",
			"City":   "Banjarnegara",
		},
	})
}
