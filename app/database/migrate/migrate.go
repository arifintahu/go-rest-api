package migrate

import (
	"github.com/arifintahu/go-rest-api/app/models"
	"gorm.io/gorm"
)

var books = []models.Book{
	{
		Title:     "Ada Apa Dengan Dunia",
		Author:    "M. Danial",
		Page:      125,
		Publisher: "Gramedia",
		Quantity:  2,
	},
	{
		Title:     "Ensiklopedia",
		Author:    "Alfonso D Alberqueque",
		Page:      439,
		Publisher: "Mizan",
		Quantity:  1,
	},
}

func Seed(db *gorm.DB) error {
	for i, _ := range books {
		err := db.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func Load(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}

	return nil
}
