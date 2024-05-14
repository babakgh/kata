package shapes

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

type Hexagon struct {
	Origin core.Point
	Side   int
}

func (c Hexagon) Draw() {
}
