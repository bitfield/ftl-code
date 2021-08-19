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

func TestBuy(t *testing.T) {
	input := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 2,
	}
	want := 1
	result := bookstore.Buy(input)
	got := result.Copies
	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of 2, got %d", want, got)
	}
}