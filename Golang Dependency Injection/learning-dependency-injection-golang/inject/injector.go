//go:build wireinject
// +build wireinject

package inject

import (
	"learning-dependency-injection-golang/app"
	"learning-dependency-injection-golang/controller"
	"learning-dependency-injection-golang/repository"
	"learning-dependency-injection-golang/service"
	"net/http"

	"github.com/google/wire"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new((*controller.CategoryControllerImpl))),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDatabase,
		NewValidator,
		categorySet,
		app.NewRouter,
		NewMiddleware,
		NewServer,
	)
	return nil
}
