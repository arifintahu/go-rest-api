package dto

type BaseResponse struct {
	Success      bool        `json:"success"`
	MessageTitle string      `json:"message_title"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	Total        int64       `json:"total,omitempty"`
}
