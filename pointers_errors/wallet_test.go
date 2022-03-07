package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(3))
		assertBalance(t, wallet, Bitcoin(3))

	})

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(3))

		assertBalance(t, wallet, Bitcoin(17))

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(200))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
