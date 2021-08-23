package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID int
}

var Books []Book

func GetBook(ID int) Book {
	for _, b := range Books {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}