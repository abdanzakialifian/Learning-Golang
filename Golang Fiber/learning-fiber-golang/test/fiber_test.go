package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "embed"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engine = mustache.New("template", ".mustache")

var app = fiber.New(fiber.Config{
	Views: engine,
	ErrorHandler: func(ctx fiber.Ctx, err error) error {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error : " + err.Error())
	},
})

func TestRouting(t *testing.T) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello Fiber")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Fiber", string(bytes))
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(ctx fiber.Ctx) error {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello " + name)
	})

	request := httptest.NewRequest(http.MethodGet, "/hello?name=Zaki", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Zaki", string(bytes))

	request = httptest.NewRequest(http.MethodGet, "/hello", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest", string(bytes))

	app.Get("/hello/v2", func(ctx fiber.Ctx) error {
		firstName := ctx.Queries()["first_name"]
		lastName := ctx.Queries()["last_name"]
		return ctx.SendString("Hello " + firstName + " " + lastName)
	})

	request = httptest.NewRequest(http.MethodGet, "/hello/v2?first_name=Abdan%20Zaki&last_name=Alifian", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Abdan Zaki Alifian", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(ctx fiber.Ctx) error {
		name := ctx.Get("name")
		cookie := ctx.Cookies("cookie")
		return ctx.SendString("Hello " + name + " " + "with cookie " + cookie)
	})

	request := httptest.NewRequest(http.MethodGet, "/request", nil)
	request.Header.Set("name", "Zaki")
	request.AddCookie(&http.Cookie{Name: "cookie", Value: "11223344"})

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Zaki with cookie 11223344", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(ctx fiber.Ctx) error {
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")
		return ctx.SendString("Get Order " + orderId + " " + "From User " + userId)
	})

	request := httptest.NewRequest(http.MethodGet, "/users/Zaki/orders/10", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 10 From User Zaki", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(ctx fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Zaki")
	request := httptest.NewRequest(http.MethodPost, "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Zaki", string(bytes))
}

//go:embed source/example.txt
var exampleFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(ctx fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}

		err = ctx.SaveFile(file, "target/"+file.Filename)
		if err != nil {
			return err
		}

		return ctx.SendString("Upload Success")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "new_example.txt")
	assert.Nil(t, err)
	file.Write(exampleFile)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	byteFiles, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(byteFiles))
}

func TestRequestBody(t *testing.T) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	app.Post("/login", func(ctx fiber.Ctx) error {
		body := ctx.Body()

		request := new(LoginRequest)

		if err := json.Unmarshal(body, request); err != nil {
			return err
		}

		return ctx.SendString("Hello " + request.Username)
	})
	body := strings.NewReader(`{"username":"Zaki", "password":"secret"}`)
	request := httptest.NewRequest(http.MethodPost, "/login", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Zaki", string(bytes))
}

func TestBodyParser(t *testing.T) {
	type RegisterRequest struct {
		Username string `json:"username" xml:"username" form:"username"`
		Password string `json:"password" xml:"password" form:"password"`
		Name     string `json:"name" xml:"name" form:"name"`
	}

	app.Post("/register", func(ctx fiber.Ctx) error {
		request := new(RegisterRequest)

		if err := ctx.Bind().Body(request); err != nil {
			return err
		}

		return ctx.SendString("Register Success " + request.Username)
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username":"Zaki", "password":"secret", "name":"Abdan Zaki Alifian"}`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Zaki", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=Zaki&password=secret&name=Abdan+Zaki+Alifian`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Zaki", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(
		`<RegisterRequest>
			<username>Zaki</username>
			<password>secret</password>
			<name>Abdan Zaki Alifian</name>
		</RegisterRequest>`,
	)

	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/xml")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Zaki", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	type User struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}

	app.Get("/user", func(ctx fiber.Ctx) error {
		return ctx.JSON(User{
			Username: "abdanzakialifian",
			Name:     "Abdan Zaki Alifian",
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"username":"abdanzakialifian","name":"Abdan Zaki Alifian"}`, string(bytes))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(ctx fiber.Ctx) error {
		return ctx.Download("source/example.txt", "new_example.txt")
	})

	request := httptest.NewRequest(http.MethodGet, "/download", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, `attachment; filename="new_example.txt"`, response.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is sample file for upload", string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	hello := func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello")
	}

	world := func(ctx fiber.Ctx) error {
		return ctx.SendString("World")
	}

	api := app.Group("/api")
	api.Get("/hello", hello) // /api/hello
	api.Get("/world", world) // /api/world

	web := app.Group("/web")
	web.Get("/hello", hello) // /web/hello
	web.Get("/world", world) // /web/world

	request := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello", string(bytes))
}

func TestStatic(t *testing.T) {
	app.Use("/public", static.New("source"))

	request := httptest.NewRequest(http.MethodGet, "/public/example.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is sample file for upload", string(bytes))
}

func TestErrorHandling(t *testing.T) {
	app.Get("/error", func(ctx fiber.Ctx) error {
		return errors.New("Internal Server Error")
	})

	request := httptest.NewRequest(http.MethodGet, "/error", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error : Internal Server Error", string(bytes))
}

func TestView(t *testing.T) {
	app.Get("/view", func(ctx fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"title":   "Hello Title",
			"header":  "Hello Header",
			"content": "Hello Content",
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/view", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "Hello Title")
	assert.Contains(t, string(bytes), "Hello Header")
	assert.Contains(t, string(bytes), "Hello Content")
}

func TestClient(t *testing.T) {
	cn := client.New()
	response, err := cn.Get("https://example.com/")
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode())
	assert.Contains(t, string(response.Body()), "Example Domain")
	defer response.Close()
}
