package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"log"

	"github.com/mashingan/smapping"
)

type PromoService interface {
	Update(c dto.Promo) dto.Promo
	All() []entity.Promo
	FindByID(categoryID uint64) entity.Promo
}

type promoService struct {
	promoRepository repository.PromoRepository
}

func NewPromoService(promoRepo repository.PromoRepository) PromoService {
	return &promoService{
		promoRepository: promoRepo,
	}
}

func (service *promoService) Update(p dto.Promo) dto.Promo {
	promo := dto.Promo{}
	err := smapping.FillStruct(&promo, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.promoRepository.UpdatePromo(promo)
	return res
}

func (service *promoService) All() []entity.Promo {
	return service.promoRepository.AllPromos()
}

func (service *promoService) FindByID(promoID uint64) entity.Promo {
	return service.promoRepository.FindPromoByID(promoID)
}
