package shapes

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

type EquilateralTriangle struct {
	Origin core.Point
	Side   int
}

func (c EquilateralTriangle) Draw() {
}
