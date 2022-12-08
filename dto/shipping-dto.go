package dto

import (
	"backend/entity"
	"strings"
	"time"
)

type ShippingUpdateDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	ShippingStatus string `json:"shipping_status" form:"shipping_status"`
	Review         string `json:"review" form:"review"`
	UserID         uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
	SizeID         uint64 `json:"size_id,omitempty" form:"size_id, omitempty"`
	AddressID      uint64 `json:"address_id,omitempty" form:"address_id, omitempty"`
	PaymentID      uint64 `json:"payment_id,omitempty" form:"payment_id, omitempty"`
	AddOnID        uint64 `json:"add_on_id,omitempty" form:"add_on_id, omitempty"`
	CategoryID     uint64 `json:"category_id,omitempty" form:"category_id, omitempty"`
}

type Shipping struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	ShippingStatus string `json:"shipping_status" form:"shipping_status"`
}

type ShippingCreateDTO struct {
	ShippingStatus string `json:"shipping_status" form:"shipping_status" `
	Review         string `json:"review" form:"review"`
	UserID         uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
	SizeID         uint64 `json:"size_id,omitempty" form:"size_id, omitempty" binding:"required"`
	AddressID      uint64 `json:"address_id,omitempty" form:"address_id, omitempty" binding:"required"`
	PaymentID      uint64 `json:"payment_id,omitempty" form:"payment_id, omitempty"`
	AddOnID        uint64 `json:"add_on_id,omitempty" form:"add_on_id, omitempty"`
	CategoryID     uint64 `json:"category_id,omitempty" form:"category_id, omitempty" binding:"required"`
}

type ShippingRequestQuery struct {
	Search string `form:"s"`
	SortBy string `form:"sortBy"`
	Sort   string `form:"sort"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type Size struct {
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

type Category struct {
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

type AddOn struct {
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

type Address struct {
	ID            uint64 `json:"id"`
	RecipientName string `json:"recipient_name"`
	FullAddress   string `json:"full_address"`
}

type Payment struct {
	ID            uint64 `json:"id"`
	PaymentStatus string `json:"payment_status"`
}

type ShippingResponse struct {
	ID             uint64   `json:"id"`
	Size           Size     `json:"size"`
	Category       Category `json:"category"`
	AddOn          AddOn    `json:"add_on"`
	Address        Address  `json:"address"`
	Payment        Payment  `json:"payment"`
	ShippingStatus string   `json:"shipping_status"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func FormatShipping(shipping *entity.Shipping) ShippingResponse {
	return ShippingResponse{
		ID:             shipping.ID,
		Size:           Size{Name: shipping.Size.Name, Price: shipping.Size.Price},
		Category:       Category{Name: shipping.Category.Name, Price: shipping.Category.Price},
		AddOn:          AddOn{Name: shipping.AddOn.Name, Price: shipping.AddOn.Price},
		Address:        Address{ID: shipping.AddressID, RecipientName: shipping.Address.RecipientName, FullAddress: shipping.Address.FullAddress},
		Payment:        Payment{ID: shipping.PaymentID, PaymentStatus: shipping.Payment.PaymentStatus},
		ShippingStatus: shipping.ShippingStatus,
		CreatedAt:      shipping.CreatedAt,
		UpdatedAt:      shipping.UpdatedAt,
	}
}

func FormatShippings(shippings []*entity.Shipping) []ShippingResponse {
	formattedShippings := []ShippingResponse{}
	for _, shipping := range shippings {
		formattedBook := FormatShipping(shipping)
		formattedShippings = append(formattedShippings, formattedBook)
	}
	return formattedShippings
}

func FormatShippingsQuery(query *ShippingRequestQuery) *ShippingRequestQuery {
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	query.SortBy = strings.ToLower(query.SortBy)
	if query.SortBy == "date" {
		query.SortBy = "updated_at"
	} else if query.SortBy == "to" {
		query.SortBy = "destination_id"
	} else if query.SortBy != "amount" {
		query.SortBy = "updated_at"
	}
	query.Sort = strings.ToUpper(query.Sort)
	if query.Sort != "ASC" {
		query.Sort = "DESC"
	}

	return query
}
