package dto

type CategoryUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}

type CategoryCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}
