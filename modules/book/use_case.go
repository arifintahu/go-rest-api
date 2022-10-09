package book

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/modules/book/types"
	"github.com/arifintahu/go-rest-api/repositories"
)

type UseCase struct {
	book repositories.IBookRepository
}

type IUseCase interface {
	ListBooks() (*[]entities.Book, error)
	GetBook(id uint64) (*entities.Book, error)
	AddBook(body *dto.BookInput) (*entities.Book, error)
	UpdateBook(id uint64, body *dto.BookInput) (*entities.Book, error)
	DeleteBook(id uint64) error
}

var _ IUseCase = (*UseCase)(nil)

func (uc UseCase) ListBooks() (*[]entities.Book, error) {
	return uc.book.ListBooks()
}

func (uc UseCase) GetBook(id uint64) (*entities.Book, error) {
	return uc.book.GetBook(id)
}

func (uc UseCase) AddBook(body *dto.BookInput) (*entities.Book, error) {
	book := entities.Book{
		Title:     body.Title,
		Author:    body.Author,
		Page:      body.Page,
		Publisher: body.Publisher,
		Quantity:  body.Quantity,
	}

	return uc.book.AddBook(&book)
}

func (uc UseCase) UpdateBook(id uint64, body *dto.BookInput) (*entities.Book, error) {
	_, err := uc.book.GetBook(id)
	if err != nil {
		return &entities.Book{}, types.ErrBookNotFound
	}

	book := entities.Book{
		Title:     body.Title,
		Author:    body.Author,
		Page:      body.Page,
		Publisher: body.Publisher,
		Quantity:  body.Quantity,
	}

	return uc.book.UpdateBook(id, &book)
}

func (uc UseCase) DeleteBook(id uint64) error {
	_, err := uc.book.GetBook(id)
	if err != nil {
		return types.ErrBookNotFound
	}

	return uc.book.DeleteBook(id)
}
