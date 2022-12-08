package dto

type PromoUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name"`
	MinFee      uint64 `json:"min_fee" `
	MaxDiscount uint64 `json:"max_discount" `
	Discount    uint64 `json:"discount" `
	Quota       uint64 `json:"quota" `
	ExpDate     string `json:"exp_date,omitempty"`
}

type Promo struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name,omitempty" form:"name"`
	MinFee      uint64 `json:"min_fee,omitempty" `
	MaxDiscount uint64 `json:"max_discount,omitempty" `
	Discount    uint64 `json:"discount,omitempty" `
	Quota       uint64 `json:"quota,omitempty" `
	ExpDate     string `json:"exp_date,omitempty"`
}

type PromoCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	MinFee      uint64 `json:"min_fee" form:"min_fee" binding:"required"`
	MaxDiscount uint64 `json:"max_discount" form:"max_discount" binding:"required"`
	Discount    uint64 `json:"discount" form:"discount" binding:"required"`
	Quota       uint64 `json:"quota" form:"quota" binding:"required"`
	ExpDate     uint64 `json:"exp_date,omitempty" form:"exp_date,omitempty"`
}
