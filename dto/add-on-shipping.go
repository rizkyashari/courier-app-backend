package dto

type AddOnShippingUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	Promotion  string `json:"promotion" form:"promotion"`
	ShippingID uint64 `json:"shipping_id,omitempty" form:"shipping_id, omitempty"`
	AddOnID    uint64 `json:"add_on_id,omitempty" form:"add_on_id, omitempty"`
}

type AddOnShippingDCreateTO struct {
	Promotion  string `json:"promotion" form:"promotion"`
	ShippingID uint64 `json:"shipping_id,omitempty" form:"shipping_id, omitempty"`
	AddOnID    uint64 `json:"add_on_id,omitempty" form:"add_on_id, omitempty"`
}
