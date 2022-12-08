package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type AddressRepository interface {
	FindByUserId(id uint64) (*entity.Address, error)
	InsertAddress(a entity.Address) entity.Address
	UpdateAddress(a entity.Address) entity.Address
	DeleteAddress(a entity.Address)
	AllAddressesAdmin() []entity.Address
	AllAddresses(userID uint64) []entity.Address
	FindAddressByID(addressID uint64) entity.Address
	FindById(id uint64) (*entity.Address, error)
}

type addressConnection struct {
	connection *gorm.DB
}

func NewAddressRepository(dbConn *gorm.DB) AddressRepository {
	return &addressConnection{
		connection: dbConn,
	}
}

func (db *addressConnection) FindByUserId(id uint64) (*entity.Address, error) {
	var address *entity.Address

	err := db.connection.Where("user_id = ?", id).Find(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (db *addressConnection) FindById(id uint64) (*entity.Address, error) {
	var address *entity.Address

	err := db.connection.Where("id = ?", id).Preload("User").Find(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}
func (db *addressConnection) InsertAddress(a entity.Address) entity.Address {
	db.connection.Save(&a)
	db.connection.Preload("User").Find(&a)
	return a
}

func (db *addressConnection) UpdateAddress(a entity.Address) entity.Address {
	db.connection.Save(&a)
	db.connection.Preload("User").Find(&a)
	return a
}

func (db *addressConnection) DeleteAddress(a entity.Address) {
	db.connection.Delete(&a)
}

func (db *addressConnection) FindAddressByID(addressID uint64) entity.Address {
	var address entity.Address
	db.connection.Preload("User").Find(&address, addressID)
	return address
}

func (db *addressConnection) AllAddressesAdmin() []entity.Address {
	var addresses []entity.Address
	db.connection.Preload("User").Find(&addresses)
	return addresses
}

func (db *addressConnection) AllAddresses(userID uint64) []entity.Address {
	var addresses []entity.Address
	db.connection.Where("user_id =?", userID).Preload("User").Find(&addresses)
	return addresses
}
