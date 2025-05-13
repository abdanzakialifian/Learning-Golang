package main

import (
	"learning-restful-api-golang/app"
	"learning-restful-api-golang/controller"
	"learning-restful-api-golang/helper"
	"learning-restful-api-golang/repository"
	"learning-restful-api-golang/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	database := app.NewDatabase()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, database, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
