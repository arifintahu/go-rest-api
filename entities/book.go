package entities

import (
	"time"
)

type Book struct {
	ID 			uint64 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Title     	string 	`gorm:"size:255;not null" json:"title"`
	Author    	string 	`gorm:"size:255;not null" json:"author"`
	Page      	uint16 	`gorm:"not null; default:0" json:"page"`
	Publisher 	string 	`gorm:"size:255" json:"publisher"`
	Quantity  	uint16 	`gorm:"not null; default:0" json:"quantity"`
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
