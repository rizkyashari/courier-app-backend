package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserPromo struct {
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	// UserID  uint64 `gorm:"type: int" json:"user_id"`
	// PromoID uint64 `gorm:"type: int" json:"promo_id"`
	UserID    uint64         `gorm:"not null" json:"-"`
	User      User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	PromoID   uint64         `gorm:"default: null" json:"-"`
	Promo     Promo          `gorm:"foreignKey:PromoID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"promo"`
	Status    int            `gorm:"type:int" json:"status"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
