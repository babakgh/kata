package core

import (
	"errors"
	"time"
)

type AccountWithdraw struct {
	account *Account
}

func NewAccountWithdraw(a *Account) AccountWithdraw {
	ad := AccountWithdraw{}
	ad.account = a
	return ad
}

func (ad *AccountWithdraw) Withdraw(date time.Time, amount int64) error {
	if ad.account.CheckBalance() < amount {
		return errors.New("IsNotAllowed")
	}

	ad.account.Transact(date, -1*amount)

	return nil
}
