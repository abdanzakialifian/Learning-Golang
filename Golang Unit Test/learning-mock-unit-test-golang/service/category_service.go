package service

import "learning-mock-unit-test-golang/entity"

type CategoryService interface {
	Get(id string) (*entity.Category, error)
}
