package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin a
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores money
type Wallet struct {
	balance Bitcoin
}

// Deposit money
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var errInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw money
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
