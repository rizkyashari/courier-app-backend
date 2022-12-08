package dto

import "backend/entity"

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type RegisterRequestBody struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestBody struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=5"`
}

type ForgotPasswordResponseBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponseBody struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	WalletNumber string `json:"wallet"`
	Token        string `json:"token"`
}

func FormatLogin(user *entity.User, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
