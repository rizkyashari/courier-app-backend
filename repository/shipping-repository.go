package repository

import (
	"backend/dto"
	"backend/entity"

	"gorm.io/gorm"
)

type ShippingRepository interface {
	InsertShipping(s entity.Shipping) entity.Shipping
	UpdateShipping(s entity.Shipping) entity.Shipping
	DeleteShipping(s entity.Shipping)
	AllShippingsAdmin() []entity.Shipping
	AllShippings(userID uint64, query *dto.ShippingRequestQuery) []*entity.Shipping
	Count(userID uint64) (int64, error)
	FindShippingByID(shippingID uint64) entity.Shipping
	FindById(id uint64) (*entity.Shipping, error)
}

type shippingConnection struct {
	connection *gorm.DB
}

func NewShippingRepository(dbConn *gorm.DB) ShippingRepository {
	return &shippingConnection{
		connection: dbConn,
	}
}

func (db *shippingConnection) InsertShipping(s entity.Shipping) entity.Shipping {
	db.connection.Save(&s)
	db.connection.Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&s)
	return s
}

func (db *shippingConnection) UpdateShipping(s entity.Shipping) entity.Shipping {
	db.connection.Save(&s)
	db.connection.Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&s)
	return s
}

func (db *shippingConnection) DeleteShipping(s entity.Shipping) {
	db.connection.Delete(&s)
}

func (db *shippingConnection) FindShippingByID(shippingID uint64) entity.Shipping {
	var shipping entity.Shipping
	db.connection.Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&shipping, shippingID)
	return shipping
}

func (db *shippingConnection) AllShippingsAdmin() []entity.Shipping {
	var shippings []entity.Shipping
	db.connection.Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&shippings)
	return shippings
}

func (db *shippingConnection) AllShippings(userID uint64, query *dto.ShippingRequestQuery) []*entity.Shipping {
	// var userID *entity.User

	var shippings []*entity.Shipping

	offset := (query.Page - 1) * query.Limit
	orderBy := query.SortBy + " " + query.Sort
	queryBuilder := db.connection.Limit(query.Limit).Offset(offset).Order(orderBy)
	queryBuilder.Debug().Where("user_id = ?", userID).Where("shipping_status ILIKE ?", "%"+query.Search+"%").Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&shippings)

	// db.connection.Where("user_id =?", userID).Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&shippings)
	return shippings
}

func (db *shippingConnection) FindById(id uint64) (*entity.Shipping, error) {
	var shipping *entity.Shipping

	err := db.connection.Where("id =?", id).Find(&shipping).Error
	if err != nil {
		return shipping, err
	}

	return shipping, nil
}

func (db *shippingConnection) Count(userID uint64) (int64, error) {
	var total int64
	database := db.connection.Model(&entity.Shipping{}).Where("user_id = ?", userID).Count(&total)

	if database.Error != nil {
		return 0, database.Error
	}

	return total, nil
}
