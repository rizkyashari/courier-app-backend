package service

import (
	"backend/entity"
	"backend/repository"
)

type AddOnService interface {
	All() []entity.AddOn
}

type addOnService struct {
	addOnRepository repository.AddOnRepository
}

func NewAddOnService(addOnRepo repository.AddOnRepository) AddOnService {
	return &addOnService{
		addOnRepository: addOnRepo,
	}
}

func (service *addOnService) All() []entity.AddOn {
	return service.addOnRepository.AllAddOns()
}
