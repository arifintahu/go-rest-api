package models

import (
	"time"
)

type Role struct {
	ID 			uint16 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name     	string 	`gorm:"size:255;not null" json:"title"`
	Slug    	string 	`gorm:"size:255;not null" json:"author"`
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
