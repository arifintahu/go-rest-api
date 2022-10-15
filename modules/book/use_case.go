package book

import (
	"time"

	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/modules/book/types"
	"github.com/arifintahu/go-rest-api/repositories"
	"github.com/arifintahu/go-rest-api/utils/pagination"
)

type UseCase struct {
	book repositories.IBookRepository
}

type IUseCase interface {
	CreateBook(body *dto.BookInput) (*entities.Book, error)
	GetBooks(query *dto.BookListQuery) (*[]entities.Book, int64, error)
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

func (uc UseCase) GetBooks(query *dto.BookListQuery) (*[]entities.Book, int64, error) {
	offset, limit := pagination.OffsetAndLimit(query.Page, query.Limit)
	params :=  dto.BookListParams{
		Offset: offset,
		Limit: limit,
	}

	books, err := uc.book.GetBooks(&params)
	if err != nil {
		return &[]entities.Book{}, 0, err
	}

	total, err := uc.book.GetBooksTotal()
	if err != nil {
		return &[]entities.Book{}, 0, err
	}

	return books, total, nil
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
		UpdatedAt: time.Now(),
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
