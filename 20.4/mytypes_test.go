package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestMyBuilderEmbedded(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.WriteString("Hello, ")
	mb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := mb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", mb.String(), wantLen, gotLen)
	}
}
