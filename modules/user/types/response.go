package types

type UserResponse struct {
	ID        uint64 `json:"id"`
	RoleID    uint16 `json:"role_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
