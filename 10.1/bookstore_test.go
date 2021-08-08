package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 2,
	}
}
