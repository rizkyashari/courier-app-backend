package entity

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	ID             uint64   `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64   `gorm:"default: null" json:"-"`
	User           User     `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	SizeID         uint64   `gorm:"default: null" json:"-"`
	Size           Size     `gorm:"foreignKey:SizeID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"size"`
	AddressID      uint64   `gorm:"default: null" json:"-"`
	Address        Address  `gorm:"foreignKey:AddressID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"address"`
	PaymentID      uint64   `gorm:"default: null" json:"-"`
	Payment        Payment  `gorm:"foreignKey:PaymentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"payment"`
	CategoryID     uint64   `gorm:"default: null" json:"-"`
	Category       Category `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
	AddOnID        uint64   `gorm:"default: null" json:"-"`
	AddOn          AddOn    `gorm:"foreignKey:AddOnID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"add_on"`
	ShippingStatus string   `gorm:"default:Waiting for payment;type:varchar(255)" json:"shipping_status"`
	Review         string   `gorm:"type:text" json:"review"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
