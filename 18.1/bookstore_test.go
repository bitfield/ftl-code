package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := bookstore.NetPriceCents(b)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
