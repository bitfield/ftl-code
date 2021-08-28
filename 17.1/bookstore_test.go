package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetBook(t *testing.T) {
	bookstore.Books = map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	got, err := bookstore.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	_, err := bookstore.GetBook(999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}
