package core_test

import (
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestTransact(t *testing.T) {
	// Setup
	log.Printf("When transacted with any amount, it has a balance with same amount")
	// Arrange
	a := core.Account{}
	a.Transact(time.Now(), 2000)
	a.Transact(time.Now(), -1000)

	want := int64(1000)
	// Act
	got := a.CheckBalance()
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

// func TestReportStatement(t *testing.T) {
// 	// Setup
// 	log.Printf("When transactions applied, it reports the statement.")
// 	// Arrange
// 	a := core.Account{}
// 	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
// 	a.Deposit(date, 1000)
// 	a.Deposit(date, 1000)
// 	a.Withdraw(date, 1000)
// 	want := []core.Transaction{{date, 1000}, {date, 1000}, {date, -1000}}
// 	// Act
// 	got := a.ReportStatement()
// 	// Assert
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("incorrect, got: %v, want %v", got, want)
// 	}
// }
