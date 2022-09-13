package services

import (
	"github.com/arifintahu/go-rest-api/app/dto"
	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/arifintahu/go-rest-api/app/repositories"
	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

type IBookService interface {
	ListBooks() (*[]models.Book, error)
	GetBook(id uint64) (*models.Book, error)
	AddBook(body *dto.BookInput) error
	UpdateBook(id uint64, body *dto.BookInput) error
	DeleteBook(id uint64) error
}

func NewBookService(db *gorm.DB) IBookService {
	return &BookService{db}
}

func (service *BookService) ListBooks() (*[]models.Book, error) {
	bookRepository := repositories.NewBookRepository(service.db)
	return bookRepository.ListBooks()
}

func (service *BookService) GetBook(id uint64) (*models.Book, error) {
	bookRepository := repositories.NewBookRepository(service.db)
	return bookRepository.GetBook(id)
}

func (service *BookService) AddBook(body *dto.BookInput) error {
	bookRepository := repositories.NewBookRepository(service.db)
	book := models.Book{
		Title: body.Title,
		Author: body.Author,
		Page: body.Page,
		Publisher: body.Publisher,
		Quantity: body.Quantity,
	}

	return bookRepository.AddBook(&book)
}

func (service *BookService) UpdateBook(id uint64, body *dto.BookInput) error {
	bookRepository := repositories.NewBookRepository(service.db)
	_, err := bookRepository.GetBook(id)
	if err != nil {
		return err
	}

	book := models.Book{
		Title: body.Title,
		Author: body.Author,
		Page: body.Page,
		Publisher: body.Publisher,
		Quantity: body.Quantity,
	}

	return bookRepository.UpdateBook(id, &book)
}

func (service *BookService) DeleteBook(id uint64) error {
	bookRepository := repositories.NewBookRepository(service.db)
	_, err := bookRepository.GetBook(id)
	if err != nil {
		return err
	}

	return bookRepository.DeleteBook(id)
}
