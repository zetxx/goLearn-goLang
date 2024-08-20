package main

import (
	"errors"
	"fmt"
)

type (
	Bitcoin int64
	Wallet  struct {
		balance Bitcoin
	}
	Stringer interface {
		String() string
	}
)

var lowBalance = errors.New("LowBalance")

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(ammount Bitcoin) {
	w.balance = w.balance + ammount
}

func (w *Wallet) Withdraw(ammount Bitcoin) error {
	if w.balance < ammount {
		return lowBalance
	}
	w.balance = w.balance - ammount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
