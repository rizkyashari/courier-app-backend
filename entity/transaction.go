package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint64 `gorm:"primaryKey"`
	SourceOfFundID *uint64
	SourceOfFund   *SourceOfFund `gorm:"foreignKey:SourceOfFundID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint64
	User           User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DestinationID  uint64
	// Wallet         Wallet `gorm:"foreignKey:DestinationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Address     Address `gorm:"foreignKey:DestinationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount      int
	Description string
	Category    string
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// func (Transaction) TableName() string {
// 	return "transactions"
// }
