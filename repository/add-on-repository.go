package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type AddOnRepository interface {
	AllAddOns() []entity.AddOn
}

type addOnConnection struct {
	connection *gorm.DB
}

func NewAddOnRepository(dbConn *gorm.DB) AddOnRepository {
	return &addOnConnection{
		connection: dbConn,
	}
}

func (db *addOnConnection) AllAddOns() []entity.AddOn {
	var addOns []entity.AddOn
	db.connection.Find(&addOns)
	return addOns
}
