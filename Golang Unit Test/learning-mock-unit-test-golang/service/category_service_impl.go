package service

import (
	"errors"
	"learning-mock-unit-test-golang/entity"
	"learning-mock-unit-test-golang/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func (service *CategoryServiceImpl) Get(id string) (*entity.Category, error) {
	category := service.CategoryRepository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
