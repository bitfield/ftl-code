package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

var Books []Book

func GetAllBooks() []Book {
	return Books
}