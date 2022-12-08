package dto

type SizeUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}

type SizeCreateDTO struct {
	Name        string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}
