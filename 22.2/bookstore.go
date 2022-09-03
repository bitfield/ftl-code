package bookstore

import "fmt"

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("unknown category %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
