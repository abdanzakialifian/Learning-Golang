package main

import (
	"learning-restful-api-golang/app"
	"learning-restful-api-golang/controller"
	"learning-restful-api-golang/helper"
	"learning-restful-api-golang/middleware"
	"learning-restful-api-golang/repository"
	"learning-restful-api-golang/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	database := app.NewDatabase()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, database, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	authMiddleware := middleware.NewAuthMiddleware(router)

	errorMiddleware := middleware.NewErrorMiddleware(authMiddleware)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: errorMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
