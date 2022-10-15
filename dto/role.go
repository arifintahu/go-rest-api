package dto

type RoleInput struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type RoleListQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type RoleListParams struct {
	Offset int
	Limit  int
}
