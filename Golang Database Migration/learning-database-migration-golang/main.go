package main

import (
	"learning-database-migration-golang/app"
	"learning-database-migration-golang/controller"
	"learning-database-migration-golang/helper"
	"learning-database-migration-golang/middleware"
	"learning-database-migration-golang/repository"
	"learning-database-migration-golang/service"
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
