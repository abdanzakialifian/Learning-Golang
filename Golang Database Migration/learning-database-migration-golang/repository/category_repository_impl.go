package repository

import (
	"context"
	"database/sql"
	"errors"
	"learning-database-migration-golang/helper"
	"learning-database-migration-golang/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, categoryName string) domain.Category {
	sql := "insert into category(name) values(?)"

	result, err := tx.ExecContext(ctx, sql, categoryName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return domain.Category{
		Id:   int(id),
		Name: categoryName,
	}
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "select * from category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "select * from category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, categoryId int, categoryName string) domain.Category {
	sql := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, categoryName, categoryId)
	helper.PanicIfError(err)

	return domain.Category{
		Id:   categoryId,
		Name: categoryName,
	}
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	sql := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
}
