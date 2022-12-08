package entity

import (
	"time"

	"gorm.io/gorm"
)

// User represents users table in database
type User struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Email       string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password    string `gorm:"->;<-;not null" json:"-"`
	Token       string `gorm:"-" json:"token,omitempty"`
	PhoneNumber uint64 `gorm:"type:int" json:"phone_number"`
	Role        string `gorm:"type:varchar(255)" json:"role"`

	ReferralCode string         `gorm:"type:text" json:"referral_code"`
	Balance      int            `gorm:"type:int" json:"balance"`
	Photos       string         `gorm:"type:text" json:"photos"`
	Addresses    *[]Address     `json:"addresses,omitempty"`
	Shippings    *[]Shipping    `json:"shippings,omitempty"`
	Payments     *[]Payment     `json:"payments,omitempty"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type PasswordReset struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token     string
	ExpiredAt time.Time
}
