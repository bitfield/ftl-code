package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID int
}

var Books map[int]Book

func GetAllBooks() []Book {
	result := []Book{}
	for _, b := range Books {
		result = append(result, b)
	}
	return result
}
