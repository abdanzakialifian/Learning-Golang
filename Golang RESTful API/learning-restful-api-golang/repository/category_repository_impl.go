package repository

import (
	"context"
	"database/sql"
	"errors"
	"learning-restful-api-golang/helper"
	"learning-restful-api-golang/model/response"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, categoryName string) response.CategoryResponse {
	sql := "insert into category(name) values(?)"

	result, err := tx.ExecContext(ctx, sql, categoryName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return response.CategoryResponse{
		Id:   int(id),
		Name: categoryName,
	}
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (response.CategoryResponse, error) {
	sql := "select * from category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := response.CategoryResponse{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []response.CategoryResponse {
	sql := "select * from category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []response.CategoryResponse{}

	for rows.Next() {
		category := response.CategoryResponse{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, categoryId int, categoryName string) response.CategoryResponse {
	sql := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, categoryName, categoryId)
	helper.PanicIfError(err)

	return response.CategoryResponse{
		Id:   categoryId,
		Name: categoryName,
	}
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	sql := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
}
