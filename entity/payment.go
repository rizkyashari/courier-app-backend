package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint64         `gorm:"primary_key:auto_increment" json:"id"`
	UserID        uint64         `gorm:"not null" json:"-"`
	User          User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	PaymentStatus string         `gorm:"type:varchar(255)" json:"payment_status"`
	TotalCost     uint64         `gorm:"type:int" json:"total_cost"`
	PromoID       uint64         `gorm:"default: null" json:"-"`
	Promo         Promo          `gorm:"foreignKey:PromoID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"promo"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
