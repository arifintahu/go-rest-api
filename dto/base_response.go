package dto

type BaseResponse struct {
	Success      bool        `json:"success"`
	MessageTitle string      `json:"message_title"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	Total        uint        `json:"total,omitempty"`
}
