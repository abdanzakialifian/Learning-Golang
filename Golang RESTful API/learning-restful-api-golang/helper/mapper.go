package helper

import (
	"learning-restful-api-golang/model/domain"
	"learning-restful-api-golang/model/response"
)

func MapCategoryResponseToCategory(response response.CategoryResponse) domain.Category {
	return domain.Category{
		Id:   response.Id,
		Name: response.Name,
	}
}

func MapCategoriesResponseToCategories(responses []response.CategoryResponse) []domain.Category {
	var categories []domain.Category
	for _, response := range responses {
		categories = append(categories, MapCategoryResponseToCategory(response))
	}
	return categories
}
