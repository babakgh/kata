package core

import (
	"errors"
	"time"
)

type Account struct {
	balance      int64
	transactions []Transaction
}

func (a *Account) History() []Transaction {
	return a.transactions
}

func (a *Account) CheckBalance() int64 {
	return a.balance
}

func (a *Account) Transact(date time.Time, amount int64) error {
	// This is a business rule
	if a.balance+amount < 0 {
		return errors.New("IsNotAllowed")
	}

	a.balance += amount
	t := NewTransaction(date, amount, a.balance)
	a.transactions = append(a.transactions, t)

	return nil
}
