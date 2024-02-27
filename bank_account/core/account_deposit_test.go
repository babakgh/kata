package core_test

import (
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestAccountDeposit(t *testing.T) {
	// Setup
	log.Printf("when deposited, it updates the account's balance")
	// Arrange
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")

	a := core.Account{}

	ad := core.NewAccountDeposit(&a)
	ad.Deposit(date, 1000)
	ad.Deposit(date, 1000)

	want := int64(2000)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
