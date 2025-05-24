package helper

import (
	"learning-database-migration-golang/model/domain"
	"learning-database-migration-golang/model/response"
)

func MapCategoryToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func MapCategoriesToCategoriesResponse(categories []domain.Category) []response.CategoryResponse {
	var responses []response.CategoryResponse
	for _, response := range categories {
		responses = append(responses, MapCategoryToCategoryResponse(response))
	}
	return responses
}
