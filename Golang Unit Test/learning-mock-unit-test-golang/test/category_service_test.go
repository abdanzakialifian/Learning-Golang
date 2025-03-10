package test

import (
	"learning-mock-unit-test-golang/entity"
	"learning-mock-unit-test-golang/service"
	"learning-mock-unit-test-golang/test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepositoryMock = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = service.CategoryServiceImpl{CategoryRepository: categoryRepositoryMock}

func TestCategoryService(t *testing.T) {
	t.Run("GetError", func(t *testing.T) {
		categoryRepositoryMock.Mock.On("FindById", "1").Return(nil)
		category, err := categoryService.Get("1")
		assert.Nil(t, category)
		assert.NotNil(t, err)
	})

	t.Run("GetSuccess", func(t *testing.T) {
		categoryEntity := entity.Category{
			Id:   "2",
			Name: "Laptop",
		}
		categoryRepositoryMock.Mock.On("FindById", "2").Return(categoryEntity)
		category, err := categoryService.Get("2")
		assert.Nil(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, categoryEntity.Id, category.Id)
		assert.Equal(t, categoryEntity.Name, category.Name)
	})
}
