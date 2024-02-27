package core

import (
	"time"
)

type RequestWithdraw struct {
	Date   time.Time
	Amount int64
}

func (rw RequestWithdraw) Withdraw(from *Account) error {
	return from.Transact(rw.Date, -1*rw.Amount)
}
