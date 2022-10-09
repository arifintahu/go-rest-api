package dto

type UserInput struct {
	RoleID    uint16 `json:"role_id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserUpdate struct {
	RoleID    uint16 `json:"role_id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
}
