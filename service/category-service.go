package service

import (
	"backend/entity"
	"backend/repository"
)

type CategoryService interface {
	All() []entity.Category
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
	}
}

func (service *categoryService) All() []entity.Category {
	return service.categoryRepository.AllCategories()
}
