package core

import "time"

type RequestDeposit struct {
	Date   time.Time
	Amount int64
}

func (rd RequestDeposit) Deposit(to *Account) error {
	return to.Transact(rd.Date, rd.Amount)
}
