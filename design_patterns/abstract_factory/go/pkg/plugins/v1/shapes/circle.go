package shapes

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

type Circle struct {
	Origin core.Point
	Radius int
}

func (c Circle) Draw() {
}
