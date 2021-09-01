package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestSetPriceCents(t *testing.T) {
	input := bookstore.Book{
		Title: "For the Love of Go",
		PriceCents: 4000,
	}
	want := 3000
	err := input.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := input.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	input := bookstore.Book{
		Title: "For the Love of Go",
		PriceCents: 4000,
	}
	err := input.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}
