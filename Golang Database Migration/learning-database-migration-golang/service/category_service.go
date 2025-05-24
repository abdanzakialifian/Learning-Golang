package service

import (
	"context"
	"learning-database-migration-golang/model/request"
	"learning-database-migration-golang/model/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
