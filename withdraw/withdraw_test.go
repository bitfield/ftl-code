package withdraw_test

import (
	"testing"
	"withdraw"
)

func TestWithdrawValid(t *testing.T) {
	t.Parallel()
	balance := 100
	amount := 20
	want := 80
	got, err := withdraw.Withdraw(balance, amount)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithdrawInvalid(t *testing.T) {
	t.Parallel()
	balance := 20
	amount := 100
	_, err := withdraw.Withdraw(balance, amount)
	if err == nil {
		t.Fatal("want error for invalid withdrawal, got nil")
	}
}
