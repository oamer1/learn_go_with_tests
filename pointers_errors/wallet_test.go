package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(3)

	got := wallet.Balance()
	// address of balance in test is 0xc000018190
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}