package exception

import (
	"learning-dependency-injection-golang/helper"
	"learning-dependency-injection-golang/model/base"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if unauthorizedError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func unauthorizedError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.WriteHeader(http.StatusUnauthorized)
		response := base.BaseResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)
		response := base.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.WriteHeader(http.StatusNotFound)
		response := base.BaseResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err any) {
	writer.WriteHeader(http.StatusInternalServerError)
	response := base.BaseResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}
	helper.WriteToResponseBody(writer, response)
}
