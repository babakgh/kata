package core_test

import (
	"errors"
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

	core.RequestDeposit{time.Now(), 1000}.Deposit(&a)
	core.RequestWithdraw{time.Now(), 1000}.Withdraw(&a)

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
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	a := core.Account{}

	rw := core.RequestWithdraw{date, 1000}

	want := errors.New("IsNotAllowed")
	// Act
	got := rw.Withdraw(&a)
	// Assert
	if got.Error() != want.Error() {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
