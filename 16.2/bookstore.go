package bookstore

import "fmt"

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID int
}

var Books map[int]Book

func GetBook(ID int) (Book, error) {
	b, ok := Books[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}