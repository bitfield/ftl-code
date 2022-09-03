package bookstore

import "errors"

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

// Buy takes a Book, decrements its number of copies, and returns the modified
// book.
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
