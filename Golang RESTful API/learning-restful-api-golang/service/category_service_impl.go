package service

import (
	"context"
	"database/sql"
	"learning-restful-api-golang/helper"
	"learning-restful-api-golang/model/domain"
	"learning-restful-api-golang/model/request"
	"learning-restful-api-golang/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) domain.Category {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response := service.CategoryRepository.Save(ctx, tx, request.Name)

	return helper.MapCategoryResponseToCategory(response)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) domain.Category {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.MapCategoryResponseToCategory(response)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []domain.Category {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	responses := service.CategoryRepository.FindAll(ctx, tx)

	return helper.MapCategoriesResponseToCategories(responses)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) domain.Category {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	response = service.CategoryRepository.Update(ctx, tx, response.Id, request.Name)

	return helper.MapCategoryResponseToCategory(response)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, response.Id)
}
