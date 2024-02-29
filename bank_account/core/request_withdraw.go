package core

import "time"

type RequestWithdraw struct {
	Date           time.Time
	Amount         int64
	FromTransacter Transacter
}

func (rw RequestWithdraw) Withdraw() error {
	return rw.FromTransacter.Transact(rw.Date, -1*rw.Amount)
}
