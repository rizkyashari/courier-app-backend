package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type AddressService interface {
	Insert(a dto.AddressCreateDTO) entity.Address
	Update(a dto.AddressUpdateDTO) entity.Address
	Delete(a entity.Address)
	All() []entity.Address
	AllByUserID(userID uint64) []entity.Address
	FindByID(addressID uint64) entity.Address
	IsAllowedToEdit(userID string, addressID uint64) bool
}

type addressService struct {
	addressRepository repository.AddressRepository
}

func NewAddressService(addressRepo repository.AddressRepository) AddressService {
	return &addressService{
		addressRepository: addressRepo,
	}
}

func (service *addressService) Insert(a dto.AddressCreateDTO) entity.Address {
	address := entity.Address{}
	err := smapping.FillStruct(&address, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.addressRepository.InsertAddress(address)
	return res
}

func (service *addressService) Update(a dto.AddressUpdateDTO) entity.Address {
	address := entity.Address{}
	err := smapping.FillStruct(&address, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.addressRepository.UpdateAddress(address)
	return res
}

func (service *addressService) Delete(a entity.Address) {
	service.addressRepository.DeleteAddress(a)
}

func (service *addressService) All() []entity.Address {
	return service.addressRepository.AllAddressesAdmin()
}

func (service *addressService) AllByUserID(userID uint64) []entity.Address {
	return service.addressRepository.AllAddresses(userID)
}

func (service *addressService) FindByID(addressID uint64) entity.Address {
	return service.addressRepository.FindAddressByID(addressID)
}

func (service *addressService) IsAllowedToEdit(userID string, addressID uint64) bool {
	b := service.addressRepository.FindAddressByID(addressID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
