package repositories

import (
	"time"

	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

type IBookRepository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	GetBooks(params *dto.BookListParams) (*[]entities.Book, error)
	GetBookDetail(id uint64) (*entities.Book, error)
	UpdateBook(id uint64, bookUpdate *entities.Book) (*entities.Book, error)
	DeleteBook(id uint64) (error)
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{db}
}

func (repo *BookRepository) CreateBook(book *entities.Book) (*entities.Book, error) {
	err := repo.db.
			Create(book).
			Take(&book).
			Error

	return book, err
}

func (repo *BookRepository) GetBooks(params *dto.BookListParams) (*[]entities.Book, error) {
	books := []entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Offset(params.Offset).
			Limit(params.Limit).
			Find(&books).
			Error

	return &books, err
}

func (repo *BookRepository) GetBookDetail(ID uint64) (*entities.Book, error) {
	book := entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Error

	return &book, err
}

func (repo *BookRepository) UpdateBook(ID uint64, bookUpdate *entities.Book) (*entities.Book, error) {
	book := entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Where("id = ?", ID).
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
			Take(&book).
			Error

	return &book, err
}

func (repo *BookRepository) DeleteBook(ID uint64) (error) {
	book := entities.Book{}
	err := repo.db.
			Model(&entities.Book{}).
			Where("id = ?", ID).
			Take(&book).
			Delete(&book).
			Error

	return err
}
