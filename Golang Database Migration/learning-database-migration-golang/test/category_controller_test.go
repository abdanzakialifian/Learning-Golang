package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"learning-database-migration-golang/app"
	"learning-database-migration-golang/controller"
	"learning-database-migration-golang/helper"
	"learning-database-migration-golang/middleware"
	"learning-database-migration-golang/repository"
	"learning-database-migration-golang/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var database = setupTestDatabase()

func setupTestDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/learning_restful_api_test")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func setupTestRouter(database *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, database, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	authMiddleware := middleware.NewAuthMiddleware(router)

	errorMiddleware := middleware.NewErrorMiddleware(authMiddleware)

	return errorMiddleware
}

func truncateCategory(db *sql.DB) {
	db.Exec("truncate category")
}

func TestCreateCategorySuccess(t *testing.T) {
	truncateCategory(database)

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(mapResponse["code"].(float64)))
	assert.Equal(t, "OK", mapResponse["status"])
	assert.Equal(t, "Gadget", mapResponse["data"].(map[string]any)["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	truncateCategory(database)

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(mapResponse["code"].(float64)))
	assert.Equal(t, "Bad Request", mapResponse["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	truncateCategory(database)

	tx, err := database.Begin()
	if err != nil {
		panic(err)
	}

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, "Gadget")
	tx.Commit()

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(mapResponse["code"].(float64)))
	assert.Equal(t, "OK", mapResponse["status"])
	assert.Equal(t, category.Id, int(mapResponse["data"].(map[string]any)["id"].(float64)))
	assert.Equal(t, category.Name, mapResponse["data"].(map[string]any)["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	truncateCategory(database)

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(mapResponse["code"].(float64)))
	assert.Equal(t, "Not Found", mapResponse["status"])
}

func TestGetCategoriesSuccess(t *testing.T) {
	truncateCategory(database)

	tx, err := database.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, "Gadget")
	category2 := categoryRepository.Save(context.Background(), tx, "Laptop")
	category3 := categoryRepository.Save(context.Background(), tx, "PC")
	tx.Commit()

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(mapResponse["code"].(float64)))

	categories := mapResponse["data"].([]any)

	categoriesResponse := categories[0].(map[string]any)
	categoriesResponse2 := categories[1].(map[string]any)
	categoriesResponse3 := categories[2].(map[string]any)

	assert.Equal(t, category.Id, int(categoriesResponse["id"].(float64)))
	assert.Equal(t, category.Name, categoriesResponse["name"])

	assert.Equal(t, category2.Id, int(categoriesResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoriesResponse2["name"])

	assert.Equal(t, category3.Id, int(categoriesResponse3["id"].(float64)))
	assert.Equal(t, category3.Name, categoriesResponse3["name"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	truncateCategory(database)

	tx, err := database.Begin()
	if err != nil {
		panic(err)
	}

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, "Gadget")
	tx.Commit()

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	requestBody := strings.NewReader(`{"name": "Laptop"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(mapResponse["code"].(float64)))
	assert.Equal(t, "OK", mapResponse["status"])
	assert.Equal(t, category.Id, int(mapResponse["data"].(map[string]any)["id"].(float64)))
	assert.Equal(t, "Laptop", mapResponse["data"].(map[string]any)["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	truncateCategory(database)

	tx, err := database.Begin()
	if err != nil {
		panic(err)
	}

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, "Gadget")
	tx.Commit()

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(mapResponse["code"].(float64)))
	assert.Equal(t, "Bad Request", mapResponse["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	truncateCategory(database)

	tx, err := database.Begin()
	if err != nil {
		panic(err)
	}

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, "Gadget")
	tx.Commit()

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(mapResponse["code"].(float64)))
	assert.Equal(t, "OK", mapResponse["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	truncateCategory(database)

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "secret")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, 404, int(mapResponse["code"].(float64)))
	assert.Equal(t, "Not Found", mapResponse["status"])
}

func TestUnauthorized(t *testing.T) {
	truncateCategory(database)

	router := setupTestRouter(database)

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var mapResponse map[string]any
	json.Unmarshal(body, &mapResponse)

	assert.Equal(t, 401, response.StatusCode)
	assert.Equal(t, 401, int(mapResponse["code"].(float64)))
	assert.Equal(t, "Unauthorized", mapResponse["status"])
}
