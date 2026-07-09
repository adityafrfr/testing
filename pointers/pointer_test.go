package pointers

import (
	"errors"
	"testing"
)

// Helpers (moved outside TestWallet)
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if !errors.Is(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func TestStringer(t *testing.T) {
	t.Run("check stringer", func(t *testing.T) {
		got := Bitcoin(10).String()
		want := "10 BTC"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
