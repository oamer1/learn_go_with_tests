package main

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
// In Go, when you call a function or a method the arguments are copied.
type Wallet struct {
	balance int
}

// In Go no need for dereferencing
func (w *Wallet) Deposit(amount int) {
	// Same as (*w).balance
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
