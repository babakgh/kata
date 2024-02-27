package core

import (
	"time"
)

type Transaction struct {
	date            time.Time
	amount          int64
	balance         int64
	transactionType string
}

func NewTransaction(date time.Time, amount int64, balance int64) Transaction {
	t := new(Transaction)
	t.date = date
	t.amount = amount
	t.balance = balance
	return *t
}
