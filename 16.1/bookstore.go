package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID int
}

var Books map[int]Book

func GetBook(ID int) Book {
	return Books[ID]
}