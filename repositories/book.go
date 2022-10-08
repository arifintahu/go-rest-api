package repositories

import (
	"time"

	"github.com/arifintahu/go-rest-api/entities"
	"github.com/palantir/stacktrace"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

type IBookRepository interface {
	ListBooks() (*[]entities.Book, error)
	GetBook(id uint64) (*entities.Book, error)
	AddBook(book *entities.Book) (error)
	UpdateBook(id uint64, bookUpdate *entities.Book) (error)
	DeleteBook(id uint64) (error)
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{db}
}

func (repo *BookRepository) ListBooks() (*[]entities.Book, error) {
	books := []entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Limit(100).
			Find(&books).
			Error

	if err != nil {
		return &[]entities.Book{}, stacktrace.NewError("Cannot get books")
	}

	return &books, nil
}

func (repo *BookRepository) GetBook(ID uint64) (*entities.Book, error) {
	book := entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Error

	if err != nil {
		return &entities.Book{}, stacktrace.NewError("Cannot find a book")
	}

	return &book, nil
}

func (repo *BookRepository) AddBook(book *entities.Book) (error) {
	err := repo.db.
			Create(book).
			Error

	if err != nil {
		return stacktrace.NewError("Cannot add new book")
	}

	return nil
}

func (repo *BookRepository) UpdateBook(ID uint64, bookUpdate *entities.Book) (error) {
	book := entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
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
	book := entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Delete(&book).
			Error

	if err != nil {
		return stacktrace.NewError("Cannot delete a book")
	}

	return nil
}
