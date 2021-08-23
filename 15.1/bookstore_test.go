package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAllBooks(t *testing.T) {
	bookstore.Books = []bookstore.Book{
		{Title: "Book 1"},
		{Title: "Book 2"},
	}
	want := []bookstore.Book{
		{Title: "Book 1"},
		{Title: "Book 2"},
	}
	got := bookstore.GetAllBooks()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
