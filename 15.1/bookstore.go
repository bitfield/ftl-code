package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}
