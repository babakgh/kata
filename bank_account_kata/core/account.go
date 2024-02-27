package core

import (
	"errors"
	"time"
)

type Account struct {
	balance      int64
	transactions []Transaction
}

// func (a *Account) ReportStatement() []Transaction {
// 	return a.transactions
// }

func (a *Account) CheckBalance() int64 {
	return a.balance
}

func (a *Account) Deposit(Date time.Time, Amount int64) error {
	a.balance += amount
	a.transactions = append(a.transactions, Transaction{Date, Amount})

	return nil
}

// func (a *Account) Withdraw(Date time.Time, Amount int64) error {
// 	if a.CheckBalance() < transaction.Amount {
// 		return errors.New("IsNotAllowed")
// 	}

// 	a.balance += -1 * amount
// 	a.transactions = append(a.transactions, Transaction{Date, -1 * Amount})

// 	return nil
// }
