package main

import (
	"learning-dependency-injection-golang/app"
	"learning-dependency-injection-golang/controller"
	"learning-dependency-injection-golang/helper"
	"learning-dependency-injection-golang/middleware"
	"learning-dependency-injection-golang/repository"
	"learning-dependency-injection-golang/service"
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
