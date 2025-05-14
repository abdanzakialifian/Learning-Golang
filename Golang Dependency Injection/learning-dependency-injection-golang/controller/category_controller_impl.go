package controller

import (
	"learning-dependency-injection-golang/helper"
	"learning-dependency-injection-golang/model/base"
	modelRequest "learning-dependency-injection-golang/model/request"
	"learning-dependency-injection-golang/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := modelRequest.CategoryCreateRequest{}

	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	category := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	response := base.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	response := base.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())

	response := base.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := modelRequest.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	category := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	response := base.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	response := base.BaseResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, response)
}
