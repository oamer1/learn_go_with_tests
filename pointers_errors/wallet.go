package main

import (
	"errors"
	"fmt"
)

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
// In Go, when you call a function or a method the arguments are copied.
// In Go, errors are values, so we can refactor it out into a variable and have a single source of truth for it.
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

// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	// errors.New creates a new error with a message of your choosing
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
