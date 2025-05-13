package service

import (
	"context"
	"learning-restful-api-golang/model/domain"
	"learning-restful-api-golang/model/request"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) domain.Category
	FindById(ctx context.Context, categoryId int) domain.Category
	FindAll(ctx context.Context) []domain.Category
	Update(ctx context.Context, request request.CategoryUpdateRequest) domain.Category
	Delete(ctx context.Context, categoryId int)
}
