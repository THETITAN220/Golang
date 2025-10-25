package main

import (
	"errors"
	"fmt"
)

type (
	Bitcoin int
	Wallet  struct {
		balance Bitcoin
	}
	Stringer interface {
		String() string
	}
)

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFund = errors.New("insufficient fund, cannot withdraw")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFund
	}
	w.balance -= amount
	return nil
}
