package core_test

import (
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestRequestWithdraw(t *testing.T) {
	// Setup
	log.Printf("when withdrawn, it updates the account's balance")
	// Arrange
	a := core.Account{}

	core.RequestDeposit{time.Now(), 1000, &a}.Deposit()
	core.RequestWithdraw{time.Now(), 1000, &a}.Withdraw()

	want := int64(0)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func TestInvalidRequestWithdraw(t *testing.T) {
	// Setup
	log.Printf("when withdrawn more than balance, it raise an error")
	// Arrange
	a := core.Account{}

	rw := core.RequestWithdraw{time.Now(), 1000, &a}

	want := core.ErrIsNotAllowed
	// Act
	got := rw.Withdraw()
	// Assert
	if got.Error() != want.Error() {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
