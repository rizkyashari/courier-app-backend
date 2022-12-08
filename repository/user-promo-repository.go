package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type UserPromoRepository interface {
	InsertUserPromo(u entity.UserPromo) entity.UserPromo
	UpdateUserPromo(u entity.UserPromo) entity.UserPromo
	AllUserPromosAdmin() []entity.UserPromo
	AllUserPromos(userID uint64) []entity.UserPromo
	FindUserPromoByID(userPromoID uint64) entity.UserPromo
	FindById(id uint64) (*entity.UserPromo, error)
}

type userPromoConnection struct {
	connection *gorm.DB
}

func NewUserPromoRepository(dbConn *gorm.DB) UserPromoRepository {
	return &userPromoConnection{
		connection: dbConn,
	}
}

func (db *userPromoConnection) InsertUserPromo(u entity.UserPromo) entity.UserPromo {
	db.connection.Save(&u)
	db.connection.Preload("User").Preload("Promo").Find(&u)
	return u
}

func (db *userPromoConnection) UpdateUserPromo(u entity.UserPromo) entity.UserPromo {
	db.connection.Save(&u)
	db.connection.Preload("User").Preload("Promo").Find(&u)
	return u
}

func (db *userPromoConnection) FindUserPromoByID(userPromoID uint64) entity.UserPromo {
	var userPromo entity.UserPromo
	db.connection.Preload("User").Preload("Promo").Find(&userPromo, userPromoID)
	return userPromo
}

func (db *userPromoConnection) AllUserPromosAdmin() []entity.UserPromo {
	var userPromos []entity.UserPromo
	db.connection.Preload("User").Preload("Promo").Find(&userPromos)
	return userPromos
}

func (db *userPromoConnection) AllUserPromos(userID uint64) []entity.UserPromo {
	// var userID *entity.User

	var userPromos []entity.UserPromo
	db.connection.Where("user_id =?", userID).Preload("User").Preload("Promo").Find(&userPromos)
	return userPromos
}

func (db *userPromoConnection) FindById(id uint64) (*entity.UserPromo, error) {
	var userPromo *entity.UserPromo

	err := db.connection.Where("id =?", id).Find(&userPromo).Error
	if err != nil {
		return userPromo, err
	}

	return userPromo, nil
}
