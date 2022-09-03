package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func GetBook(catalog []Book, ID int) Book {
	for _, b := range catalog {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}
