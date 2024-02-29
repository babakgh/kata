package core

import "time"

type RequestDeposit struct {
	Date         time.Time
	Amount       int64
	ToTransacter Transacter
}

func (rd RequestDeposit) Deposit() error {
	return rd.ToTransacter.Transact(rd.Date, rd.Amount)
}
