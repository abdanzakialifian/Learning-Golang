package repository

import (
	"context"
	"database/sql"
	"learning-dependency-injection-golang/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, categoryName string) domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Update(ctx context.Context, tx *sql.Tx, categoryId int, categoryName string) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
}
