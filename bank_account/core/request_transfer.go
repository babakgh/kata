package core

import "time"

type RequestTransfer struct {
	Date   time.Time
	Amount int64
}

func (rt RequestTransfer) Transfer(from *Account, to *Account) error {
	if err := from.Transact(rt.Date, -1*rt.Amount); err != nil {
		return err
	}
	return to.Transact(rt.Date, rt.Amount)
	// We cant test this atm
	// if err := to.Transact(rt.Date, rt.Amount); err != nil {
	// 	// Well something bad happend and we should refund the account above!
	// 	return err
	// }
	// return nil
}
