package core

import "time"

type Transacter interface {
	Transact(date time.Time, amount int64) error
}
