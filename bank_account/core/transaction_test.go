package core_test

import (
	"log"
	"testing"
	"time"

	"github.com/babakgh/kata/bank_account/core"
)

func TestTransaction(t *testing.T) {
	// Setup
	log.Printf("On a new transaction, it has date, amount and balance")
	// Arrange
	date, _ := time.Parse("%Y-%m-%d", "2024-01-01")
	amount := int64(1000)
	balance := int64(2000)

	want := core.Transaction{date, amount, balance}
	// Act
	got := core.NewTransaction(date, amount, balance)
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
