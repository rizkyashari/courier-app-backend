package dto

import "backend/entity"

type UserUpdateDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Name        string `json:"name" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email" `
	PhoneNumber uint64 `json:"phone_number" form:"phone_number"`
	Password    string `json:"password,omitempty" form:"password,omitempty"`
}

type User struct {
	ID          uint64 `json:"id" form:"id"`
	Name        string `json:"name" form:"name" `
	Email       string `json:"email" form:"email"  `
	PhoneNumber uint64 `json:"phone_number" form:"phone_number"`
	Photos      string `json:"photos" `
	Password    string `json:"password,omitempty" form:"password,omitempty"`
}

// type UserCreateDTO struct {
// 	ID       uint64 `json:"id" form:"id" binding:"required"`
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
// }

type UserRequestParams struct {
	UserID uint64 `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type UserResponseBody struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user *entity.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	return formattedUser
}

func FormatUsers(authors []*entity.User) []UserResponseBody {
	formattedUsers := []UserResponseBody{}
	for _, user := range authors {
		formattedUser := FormatUser(user)
		formattedUsers = append(formattedUsers, formattedUser)
	}
	return formattedUsers
}

type UserDetailResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUserDetail(user *entity.User) UserDetailResponse {
	formattedUser := UserDetailResponse{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	return formattedUser
}
