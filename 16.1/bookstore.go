package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func GetBook(catalog map[int]Book, ID int) Book {
	return catalog[ID]
}
