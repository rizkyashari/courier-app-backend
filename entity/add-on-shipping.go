package entity

type AddOnShipping struct {
	Promotion  string   `gorm:"type:text" json:"promotion"`
	ShippingID uint64   `gorm:"not null" json:"-"`
	Shipping   Shipping `gorm:"foreignKey:ShippingID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"shipping"`
	AddOnID    uint64   `gorm:"not null" json:"-"`
	AddOn      AddOn    `gorm:"foreignKey:AddOnID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"add_on"`
}
