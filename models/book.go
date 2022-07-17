package models

//Book defines table structure
type Book struct {
	ID 			int 	`json: "id"`
	Title     	string 	`json: "title"`
	Author    	string 	`json: "author"`
	Page      	uint16 	`json: "page"`
	Publisher 	string 	`json: "publisher"`
	Quantity  	uint16 	`json: "quantity"`
}

//BookStorage defines all databases operations
type BookStorage interface {
	ListBooks() ([]Book, error)
	GetBook(id int) (Book, error)
	// AddBook(book Book) (Book, error)
	// UpdateBook(book Book) (Book, error)
	// DeleteBook(id int) error
}
