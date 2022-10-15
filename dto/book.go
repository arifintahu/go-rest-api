package dto

type BookInput struct {
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Page      uint16 `json:"page" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	Quantity  uint16 `json:"quantity" binding:"required"`
}

type BookListQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type BookListParams struct {
	Offset int
	Limit  int
}
