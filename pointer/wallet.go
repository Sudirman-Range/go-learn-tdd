package main

import "errors"

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Withdraw(amount int) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
