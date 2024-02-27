package core

import "time"

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

func (a *Account) Transact(date time.Time, amount int64) {
	a.balance += amount
	t := NewTransaction(date, amount, a.balance)
	a.transactions = append(a.transactions, t)
}
