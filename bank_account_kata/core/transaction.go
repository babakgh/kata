package core

import (
	"time"
)

type Transaction struct {
	Date   time.Time
	Amount int64
	// Balance int64
}
