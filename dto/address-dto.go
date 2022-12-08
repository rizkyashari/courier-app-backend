package dto

type AddressUpdateDTO struct {
	ID                   uint64 `json:"id" form:"id" binding:"required"`
	FullAddress          string `json:"full_address" form:"full_address" binding:"required"`
	RecipientName        string `json:"recipient_name" form:"recipient_name" binding:"required"`
	RecipientPhoneNumber uint64 `json:"recipient_phone_number" form:"recipient_phone_number" binding:"required"`
	UserID               uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
}

type AddressCreateDTO struct {
	FullAddress          string `json:"full_address" form:"full_address" binding:"required"`
	RecipientName        string `json:"recipient_name" form:"recipient_name" binding:"required"`
	RecipientPhoneNumber uint64 `json:"recipient_phone_number" form:"recipient_phone_number" binding:"required"`
	UserID               uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
}
