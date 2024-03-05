package shapes

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

type Circle struct {
	Points []core.Point
}

func (c Circle) Draw() {
}
