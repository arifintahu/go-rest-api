package services

import (
	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/palantir/stacktrace"
	"gorm.io/gorm"
)

type BookService struct {
	DB *gorm.DB
}

//Apply interface BookStorage
var _ models.BookStorage = (*BookService)(nil)

var bookList = []*models.Book{
	{
		ID: 1,
		Title:     "Ada Apa Dengan Dunia",
		Author:    "M. Danial",
		Page:      125,
		Publisher: "Gramedia",
		Quantity:  2,
	},
	{
		ID: 2,
		Title:     "Ensiklopedia",
		Author:    "Alfonso D Alberqueque",
		Page:      439,
		Publisher: "Mizan",
		Quantity:  1,
	},
}

func (service *BookService) ListBooks() (*[]models.Book, error) {
	books := []models.Book{}
	err := service.DB.
			Model(&models.Book{}).
			Limit(100).
			Find(&books).
			Error

	if err != nil {
		return &[]models.Book{}, stacktrace.NewError("Cannot get books")
	}

	return &books, nil
}

func (service *BookService) GetBook(ID uint64) (*models.Book, error) {
	book := models.Book{}
	err := service.DB.
			Model(&models.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Error

	if err != nil {
		return &models.Book{}, stacktrace.NewError("Cannot find a book")
	}

	return &book, nil
}
