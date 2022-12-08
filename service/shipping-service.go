package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type ShippingService interface {
	Insert(s dto.ShippingCreateDTO) entity.Shipping
	Update(s dto.ShippingUpdateDTO) entity.Shipping
	Delete(s entity.Shipping)
	AllShippings() []entity.Shipping
	All(userID uint64, query *dto.ShippingRequestQuery) []*entity.Shipping
	FindByID(shippingID uint64) entity.Shipping
	IsAllowedToEdit(userID string, shippingID uint64) bool
	CountShipping(userID uint64) (int64, error)
}

type shippingService struct {
	shippingRepository repository.ShippingRepository
}

func NewShippingService(shippingRepo repository.ShippingRepository) ShippingService {
	return &shippingService{
		shippingRepository: shippingRepo,
	}
}

func (service *shippingService) CountShipping(userID uint64) (int64, error) {
	totalShippings, err := service.shippingRepository.Count(userID)
	if err != nil {
		return totalShippings, err
	}

	return totalShippings, nil
}
func (service *shippingService) Insert(s dto.ShippingCreateDTO) entity.Shipping {
	shipping := entity.Shipping{}
	err := smapping.FillStruct(&shipping, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.shippingRepository.InsertShipping(shipping)
	return res
}

func (service *shippingService) Update(s dto.ShippingUpdateDTO) entity.Shipping {
	shipping := entity.Shipping{}
	err := smapping.FillStruct(&shipping, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.shippingRepository.UpdateShipping(shipping)
	return res
}

func (service *shippingService) Delete(s entity.Shipping) {
	service.shippingRepository.DeleteShipping(s)
}

func (service *shippingService) All(userID uint64, query *dto.ShippingRequestQuery) []*entity.Shipping {
	return service.shippingRepository.AllShippings(userID, query)
}

func (service *shippingService) AllShippings() []entity.Shipping {
	return service.shippingRepository.AllShippingsAdmin()
}
func (service *shippingService) FindByID(shippingID uint64) entity.Shipping {
	return service.shippingRepository.FindShippingByID(shippingID)
}

func (service *shippingService) IsAllowedToEdit(userID string, shippingID uint64) bool {
	b := service.shippingRepository.FindShippingByID(shippingID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
