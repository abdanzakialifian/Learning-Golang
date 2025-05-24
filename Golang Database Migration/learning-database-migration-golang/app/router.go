package app

import (
	"learning-database-migration-golang/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/categories", controller.Create)
	router.GET("/api/categories/:id", controller.FindById)
	router.GET("/api/categories", controller.FindAll)
	router.PUT("/api/categories/:id", controller.Update)
	router.DELETE("/api/categories/:id", controller.Delete)

	return router
}
