package dto

type RegisterDTO struct {
	Name         string `json:"name" form:"name" binding:"required"`
	Email        string `json:"email" form:"email" binding:"required,email" `
	Password     string `json:"password" form:"password" binding:"required"`
	PhoneNumber  uint64 `json:"phone_number,string,omitempty"`
	ReferralCode string `json:"referral_code" `
}

type AdminRegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}
