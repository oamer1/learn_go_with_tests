package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(3))
		assertBalance(t, wallet, Bitcoin(3))

	})

	t.Run("Withdraw with funds", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		// What this is telling us is that we have not checked the error being returned on that line of code
		wallet.Withdraw(Bitcoin(3))

		assertBalance(t, wallet, Bitcoin(17))

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
