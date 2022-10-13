package dto

type AccountLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountLoginResponse struct {
	Token string `json:"token"`
}
