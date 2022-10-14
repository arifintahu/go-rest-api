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
	CreateBook(body *dto.BookInput) (*entities.Book, error)
	GetBooks() (*[]entities.Book, error)
	GetBookDetail(id uint64) (*entities.Book, error)
	UpdateBook(id uint64, body *dto.BookInput) (*entities.Book, error)
	DeleteBook(id uint64) error
}

var _ IUseCase = (*UseCase)(nil)

func (uc UseCase) CreateBook(body *dto.BookInput) (*entities.Book, error) {
	book := entities.Book{
		Title:     body.Title,
		Author:    body.Author,
		Page:      body.Page,
		Publisher: body.Publisher,
		Quantity:  body.Quantity,
	}

	return uc.book.CreateBook(&book)
}

func (uc UseCase) GetBooks() (*[]entities.Book, error) {
	return uc.book.GetBooks()
}

func (uc UseCase) GetBookDetail(id uint64) (*entities.Book, error) {
	return uc.book.GetBookDetail(id)
}

func (uc UseCase) UpdateBook(id uint64, body *dto.BookInput) (*entities.Book, error) {
	_, err := uc.book.GetBookDetail(id)
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
	_, err := uc.book.GetBookDetail(id)
	if err != nil {
		return types.ErrBookNotFound
	}

	return uc.book.DeleteBook(id)
}
