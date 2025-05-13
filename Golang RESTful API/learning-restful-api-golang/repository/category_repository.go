package repository

import (
	"context"
	"database/sql"
	"learning-restful-api-golang/model/response"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, categoryName string) response.CategoryResponse
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (response.CategoryResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []response.CategoryResponse
	Update(ctx context.Context, tx *sql.Tx, categoryId int, categoryName string) response.CategoryResponse
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
}
