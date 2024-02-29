package core

import (
	"time"
)

type Transaction struct {
	date    time.Time
	amount  int64
	balance int64
}

func NewTransaction(date time.Time, amount int64, balance int64) Transaction {
	return Transaction{date: date, amount: amount, balance: balance}
}
