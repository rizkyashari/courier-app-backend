package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name        string         `gorm:"type:varchar(255)" json:"name"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	Price       uint64         `gorm:"type:int" json:"price"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
