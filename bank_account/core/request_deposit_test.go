package core_test

import (
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestRequestDeposit(t *testing.T) {
	// Setup
	log.Printf("when deposited, it updates the account's balance")
	// Arrange
	a := core.Account{}

	core.RequestDeposit{time.Now(), 1000}.Deposit(&a)
	core.RequestDeposit{time.Now(), 2000}.Deposit(&a)

	want := int64(3000)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
