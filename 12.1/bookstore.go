package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

// Buy takes a Book, decrements its number of copies, and returns the modified
// book.
func Buy(b Book) Book {
	b.Copies--
	return b
}
