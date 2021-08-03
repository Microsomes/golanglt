package main

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balanace Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balanace += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balanace {
		return ErrInsufficientFunds
	}
	w.balanace -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balanace
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
