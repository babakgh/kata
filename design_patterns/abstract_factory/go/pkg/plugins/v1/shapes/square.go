package shapes

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

type Square struct {
	Origin core.Point
	Side   int
}

func (s Square) Draw() {
}
