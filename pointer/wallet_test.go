package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)
		if err != nil {
			t.Error(err.Error())
		}
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, 20)

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want int) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")

	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
