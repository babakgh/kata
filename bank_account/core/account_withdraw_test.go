package core_test

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestAccountWithdraw(t *testing.T) {
	// Setup
	log.Printf("when withdrawn, it updates the account's balance")
	// Arrange
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")

	a := core.Account{}

	ad := core.NewAccountDeposit(&a)
	ad.Deposit(date, 1000)

	aw := core.NewAccountWithdraw(&a)
	aw.Withdraw(date, 1000)

	want := int64(0)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func TestAccountInvalidWithdraw(t *testing.T) {
	// Setup
	log.Printf("when withdrawn more than balance, it raise an error")
	// Arrange
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	a := core.Account{}

	aw := core.NewAccountWithdraw(&a)
	aw.Withdraw(date, 1000)

	want := errors.New("IsNotAllowed")
	// Act
	got := aw.Withdraw(date, 1000)
	// Assert
	if got.Error() != want.Error() {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
