package services

import (
	"github.com/arifintahu/go-rest-api/app/models"

	"github.com/palantir/stacktrace"
)

type BookService struct {}

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

func (service *BookService) ListBooks() ([]models.Book, error) {
	var list []models.Book
	for _,v := range bookList {
		list = append(list, *v)
	}

	if len(list) == 0 {
		return []models.Book{}, stacktrace.NewError("List books not found")
	}
	return list, nil
}

func (service *BookService) GetBook(ID uint64) (models.Book, error) {
	var book models.Book
	for _, v := range bookList {
		if v.ID == ID {
			book = *v
			break
		}
	}
	if book.ID != ID {
		return models.Book{}, stacktrace.NewError("Cannot find a book")
	}
	return book, nil
}
