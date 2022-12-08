package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type SizeRepository interface {
	AllSizes() []entity.Size
}

type sizeConnection struct {
	connection *gorm.DB
}

func NewSizeRepository(dbConn *gorm.DB) SizeRepository {
	return &sizeConnection{
		connection: dbConn,
	}
}

func (db *sizeConnection) AllSizes() []entity.Size {
	var sizes []entity.Size
	db.connection.Find(&sizes)
	return sizes
}
