package main

import (
	"errors"
	"fmt"
)

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
// In Go, when you call a function or a method the arguments are copied.

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// In Go no need for dereferencing
func (w *Wallet) Deposit(amount Bitcoin) {
	// Same as (*w).balance
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	// errors.New creates a new error with a message of your choosing
	if amount > w.balance {
		return errors.New("Oh no")
	}
	w.balance -= amount
	return nil
}
