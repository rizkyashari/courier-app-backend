package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	AllCategories() []entity.Category
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(dbConn *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: dbConn,
	}
}

func (db *categoryConnection) AllCategories() []entity.Category {
	var categories []entity.Category
	db.connection.Find(&categories)
	return categories
}
