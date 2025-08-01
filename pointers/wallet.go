package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficentFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficentFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
