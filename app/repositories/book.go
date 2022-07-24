package repositories

import (
	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/palantir/stacktrace"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

//Apply interface BookStorage
var _ models.BookStorage = (*BookRepository)(nil)

func (repo *BookRepository) ListBooks() (*[]models.Book, error) {
	books := []models.Book{}
	err := repo.DB.
			Model(&models.Book{}).
			Limit(100).
			Find(&books).
			Error

	if err != nil {
		return &[]models.Book{}, stacktrace.NewError("Cannot get books")
	}

	return &books, nil
}

func (repo *BookRepository) GetBook(ID uint64) (*models.Book, error) {
	book := models.Book{}
	err := repo.DB.
			Model(&models.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Error

	if err != nil {
		return &models.Book{}, stacktrace.NewError("Cannot find a book")
	}

	return &book, nil
}
