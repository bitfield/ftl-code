package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetBook(t *testing.T) {
	bookstore.Books = []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	got := bookstore.GetBook(2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
