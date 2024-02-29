package core_test

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestRequestTransfer(t *testing.T) {
	// Setup
	log.Printf("when withdrawn, it updates the account's balance")
	// Arrange
	from := core.Account{}
	core.RequestDeposit{time.Now(), 1000, &from}.Deposit()

	to := core.Account{}

	t.Logf("%v %v", from, to)

	if err := (core.RequestTransfer{time.Now(), 300, &from, &to}.Transfer()); err != nil {
		t.Errorf("It's a bug! Fix it: %v", err)
	}

	want := [2]int64{700, 300}
	// Act
	got := [2]int64{from.CheckBalance(), to.CheckBalance()}
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func TestInvalidRequestTransfer(t *testing.T) {
	// Setup
	log.Printf("when withdrawn more than balance, it raise an error")
	// Arrange
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")

	from := core.Account{}
	to := core.Account{}
	rt := core.RequestTransfer{date, 300, &from, &to}

	want := errors.New("IsNotAllowed")
	// Act
	got := rt.Transfer()
	// Assert
	if got.Error() != want.Error() {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
