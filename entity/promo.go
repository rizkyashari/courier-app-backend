package entity

import (
	"time"

	"gorm.io/gorm"
)

type Promo struct {
	ID          uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name        string         `gorm:"type:varchar(255)" json:"name"`
	MinFee      uint64         `gorm:"type:int" json:"min_fee"`
	MaxDiscount uint64         `gorm:"type:int" json:"max_discount"`
	Discount    uint64         `gorm:"type:int" json:"discount"`
	Quota       uint64         `gorm:"type:int" json:"quota"`
	ExpDate     time.Time      `json:"exp_date"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
