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
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("with price %d, after %d%% discount want net %d, got %d", b.PriceCents, b.DiscountPercent, want, got)
	}
}
