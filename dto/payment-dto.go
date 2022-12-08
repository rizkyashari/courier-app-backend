package dto

type PaymentUpdateDTO struct {
	ID            uint64 `json:"id" form:"id" binding:"required"`
	PaymentStatus string `json:"payment_status" form:"payment_status" binding:"required"`
	TotalCost     uint64 `json:"total_cost" form:"total_cost" binding:"required"`
	PromoID       uint64 `json:"promo_id,omitempty" form:"promo_id, omitempty"`
	UserID        uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
}

type PaymentCreateDTO struct {
	PaymentStatus string `json:"payment_status" form:"payment_status" binding:"required"`
	TotalCost     uint64 `json:"total_cost" form:"total_cost" binding:"required"`
	PromoID       uint64 `json:"promo_id,omitempty" form:"promo_id, omitempty"`
	UserID        uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
}
