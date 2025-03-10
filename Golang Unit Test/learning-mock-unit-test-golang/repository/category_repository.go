package repository

import "learning-mock-unit-test-golang/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
