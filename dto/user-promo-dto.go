package dto

type UserPromoUpdateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	UserID  uint64 `json:"user_id,omitempty" form:"user_id, omitempty" binding:"required"`
	PromoID uint64 `json:"promo_id,omitempty" form:"size_id, omitempty" binding:"required"`
	Status  int    `json:"status" form:"status"`
}

type UserPromoCreateDTO struct {
	UserID  uint64 `json:"user_id,omitempty" form:"user_id, omitempty" binding:"required"`
	PromoID uint64 `json:"promo_id,omitempty" form:"size_id, omitempty" binding:"required"`
	Status  int    `json:"status" form:"status"`
}
