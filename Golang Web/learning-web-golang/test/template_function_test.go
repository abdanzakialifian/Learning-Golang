package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
}

func TestTemplateFunction(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateFunction(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateFunction(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Zaki"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Abdan",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateFunctionGlobal(writer http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Zaki",
	})
}

func TestFunctionCreateGlobal(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateFunctionCreateGlobal(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateFunctionCreateGlobal(writer http.ResponseWriter, _ *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Abdan Zaki Alifian",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	templateFunctionPipelines(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateFunctionPipelines(writer http.ResponseWriter, _ *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{.SayHello "Brother" | upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Abdan Zaki",
	})
}
