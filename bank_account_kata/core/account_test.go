package core_test

import (
	"errors"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/babakgh/bank/core"
)

func TestDeposit(t *testing.T) {
	// Setup
	log.Printf("when deposited, it updates the account's balance")
	// Arrange
	a := core.Account{}
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	a.Deposit(date, int64(1000))
	a.Deposit(date, int64(1000))
	want := int64(2000)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	// Setup
	log.Printf("when withdrawn, it updates the account's balance")
	// Arrange
	a := core.Account{}
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	a.Deposit(date, int64(1000))
	a.Withdraw(date, int64(1000))
	want := int64(0)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func TestInvalidWithdraw(t *testing.T) {
	// Setup
	log.Printf("when withdrawn more than balance, it raise an error")
	// Arrange
	a := core.Account{}
	want := errors.New("IsNotAllowed")
	// Act
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	a.Withdraw(date, int64(1000))
	// Assert
	if got.Error() != want.Error() {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func TestReportStatement(t *testing.T) {
	// Setup
	log.Printf("When transactions applied, it reports the statement.")
	// Arrange
	a := core.Account{}
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	a.Deposit(date, int64(1000))
	a.Deposit(date, int64(1000))
	a.Withdraw(date, int64(1000))
	want := []core.Transaction{{date, 1000}, {date, 1000}, {date, -1000}}
	// Act
	got := a.ReportStatement()
	// Assert
	if !reflect.DeepEqual(got, want) {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
