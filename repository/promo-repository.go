package repository

import (
	"backend/dto"
	"backend/entity"

	"gorm.io/gorm"
)

type PromoRepository interface {
	UpdatePromo(p dto.Promo) dto.Promo
	AllPromos() []entity.Promo
	FindPromoByID(promoID uint64) entity.Promo
}

type promoConnection struct {
	connection *gorm.DB
}

func NewPromoRepository(dbConn *gorm.DB) PromoRepository {
	return &promoConnection{
		connection: dbConn,
	}
}

func (db *promoConnection) UpdatePromo(p dto.Promo) dto.Promo {
	db.connection.Debug().Save(&p)
	return p
}

func (db *promoConnection) FindPromoByID(promoID uint64) entity.Promo {
	var promo entity.Promo
	db.connection.Find(&promo, promoID)
	return promo
}

func (db *promoConnection) AllPromos() []entity.Promo {
	var promos []entity.Promo
	db.connection.Find(&promos)
	return promos
}
