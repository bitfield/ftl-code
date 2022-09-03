package bookstore

import "fmt"

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}
