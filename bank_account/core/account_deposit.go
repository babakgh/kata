package core

import "time"

type AccountDeposit struct {
	account *Account
}

func NewAccountDeposit(a *Account) AccountDeposit {
	ad := AccountDeposit{}
	ad.account = a
	return ad
}

func (ad *AccountDeposit) Deposit(date time.Time, amount int64) error {
	ad.account.Transact(date, amount)

	return nil
}
