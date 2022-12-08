package service

import (
	"backend/entity"
	"backend/repository"
)

type SizeService interface {
	All() []entity.Size
}

type sizeService struct {
	sizeRepository repository.SizeRepository
}

func NewSizeService(sizeRepo repository.SizeRepository) SizeService {
	return &sizeService{
		sizeRepository: sizeRepo,
	}
}

func (service *sizeService) All() []entity.Size {
	return service.sizeRepository.AllSizes()
}
