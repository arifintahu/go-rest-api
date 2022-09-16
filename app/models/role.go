package models

import (
	"time"
)

type Role struct {
	ID 			uint16 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name     	string 	`gorm:"size:255;not null" json:"name"`
	Slug    	string 	`gorm:"size:255;not null" json:"slug"`
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
