package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type UserPromoService interface {
	Insert(s dto.UserPromoCreateDTO) entity.UserPromo
	Update(s dto.UserPromoUpdateDTO) entity.UserPromo
	AllUserPromos() []entity.UserPromo
	All(userID uint64) []entity.UserPromo
	FindByID(userPromoID uint64) entity.UserPromo
	IsAllowedToEdit(userID string, userPromoID uint64) bool
}

type userPromoService struct {
	userPromoRepository repository.UserPromoRepository
}

func NewUserPromoService(userPromoRepo repository.UserPromoRepository) UserPromoService {
	return &userPromoService{
		userPromoRepository: userPromoRepo,
	}
}

func (service *userPromoService) Insert(s dto.UserPromoCreateDTO) entity.UserPromo {
	userPromo := entity.UserPromo{}
	err := smapping.FillStruct(&userPromo, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userPromoRepository.InsertUserPromo(userPromo)
	return res
}

func (service *userPromoService) Update(s dto.UserPromoUpdateDTO) entity.UserPromo {
	userPromo := entity.UserPromo{}
	err := smapping.FillStruct(&userPromo, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userPromoRepository.UpdateUserPromo(userPromo)
	return res
}

func (service *userPromoService) All(userID uint64) []entity.UserPromo {
	return service.userPromoRepository.AllUserPromos(userID)
}

func (service *userPromoService) AllUserPromos() []entity.UserPromo {
	return service.userPromoRepository.AllUserPromosAdmin()
}
func (service *userPromoService) FindByID(userPromoID uint64) entity.UserPromo {
	return service.userPromoRepository.FindUserPromoByID(userPromoID)
}

func (service *userPromoService) IsAllowedToEdit(userID string, userPromoID uint64) bool {
	b := service.userPromoRepository.FindUserPromoByID(userPromoID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
