package entity

import (
	"time"

	"gorm.io/gorm"
)

// User represents users table in database
type Admin struct {
	ID          uint64      `gorm:"primary_key:auto_increment" json:"id"`
	Name        string      `gorm:"type:varchar(255)" json:"name"`
	Email       string      `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password    string      `gorm:"->;<-;not null" json:"-"`
	Token       string      `gorm:"-" json:"token,omitempty"`
	PhoneNumber uint64      `gorm:"type:int" json:"phone_number"`
	Role        string      `gorm:"type:varchar(255)" json:"role"`
	Photos      string      `gorm:"type:text" json:"photos"`
	Shippings   *[]Shipping `json:"shippings,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
