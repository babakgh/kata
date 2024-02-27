package core

import (
	"time"
)

type Transaction struct {
	Date    time.Time
	Amount  int64
	Balance int64
}

func NewTransaction(date time.Time, amount int64, balance int64) Transaction {
	return Transaction{date, amount, balance}
}
