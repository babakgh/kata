package core_test

import (
	"log"
	"reflect"
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

func TestHistory(t *testing.T) {
	// Setup
	log.Printf("When transactions are applied, it keeps history.")
	// Arrange
	date, _ := time.Parse(time.DateOnly, "2024-01-01")
	a := core.Account{}

	core.RequestDeposit{date, 1000}.Deposit(&a)
	core.RequestDeposit{date, 2000}.Deposit(&a)
	core.RequestWithdraw{date, 800}.Withdraw(&a)
	want := []core.Transaction{
		core.NewTransaction(date, 1000, 1000),
		core.NewTransaction(date, 2000, 3000),
		core.NewTransaction(date, -800, 2200)}
	// Act
	got := a.History()
	// Assert
	if !reflect.DeepEqual(got, want) {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
