package core

import (
	"time"
)

type RequestTransfer struct {
	Date        time.Time
	Amount      int64
	FromAccount Transacter
	ToAccount   Transacter
}

func (rt RequestTransfer) Transfer() error {
	if err := rt.FromAccount.Transact(rt.Date, -1*rt.Amount); err != nil {
		return err
	}
	if err := rt.ToAccount.Transact(rt.Date, rt.Amount); err != nil {
		// TODO: Refund to FromAccount
		return err
	}
	return nil
}
