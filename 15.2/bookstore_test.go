package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
	}
	want := bookstore.Book{ID: 1, Title: "For the Love of Go"}
	got := bookstore.GetBook(catalog, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
