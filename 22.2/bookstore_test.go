package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestSetCategory(t *testing.T) {
	input := bookstore.Book{
		Title: "For the Love of Go",
	}
	want := "Autobiography"
	err := input.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := input.Category()
	if want != got {
		t.Errorf("want category %q, got %q", want, got)
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	input := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := input.SetCategory("bogus")
	if err == nil {
		t.Fatal("want error setting invalid category 'bogus', got nil")
	}
}
