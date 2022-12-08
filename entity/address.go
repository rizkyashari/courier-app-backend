package entity

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID                   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID               uint64 `gorm:"default: null" json:"-"`
	User                 User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	FullAddress          string `gorm:"type:varchar(255)" json:"full_address"`
	RecipientName        string `gorm:"type:varchar(255)" json:"recipient_name"`
	RecipientPhoneNumber uint64 `gorm:"type:int" json:"recipient_phone_number"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
}
