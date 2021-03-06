package repositories

import (
	"time"

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

func (repo *BookRepository) AddBook(book *models.Book) (error) {
	err := repo.DB.
			Create(book).
			Error

	if err != nil {
		return stacktrace.NewError("Cannot add new book")
	}

	return nil
}

func (repo *BookRepository) UpdateBook(ID uint64, bookUpdate *models.Book) (error) {
	book := models.Book{}
	err := repo.DB.
			Model(&models.Book{}).
			Where("id = ?", ID).
			Take(&book).
			UpdateColumns(
				map[string]interface{}{
					"title": bookUpdate.Title,
					"author": bookUpdate.Author,
					"page": bookUpdate.Page,
					"publisher": bookUpdate.Publisher,
					"quantity": bookUpdate.Quantity,
					"updated_at": time.Now(),
				},
			).
			Error

	if err != nil {
		return stacktrace.NewError("Cannot update a book")
	}

	return nil
}

func (repo *BookRepository) DeleteBook(ID uint64) (error) {
	book := models.Book{}
	err := repo.DB.
			Model(&models.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Delete(&book).
			Error

	if err != nil {
		return stacktrace.NewError("Cannot delete a book")
	}

	return nil
}
