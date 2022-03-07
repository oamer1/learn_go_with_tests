package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(3))

		got := wallet.Balance()
		// address of balance in test is 0xc000018190
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		want := Bitcoin(3)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(3))

		got := wallet.Balance()
		// address of balance in test is 0xc000018190
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		want := Bitcoin(3)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
