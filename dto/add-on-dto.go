package dto

type AddOnUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}

type AddOnCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}
