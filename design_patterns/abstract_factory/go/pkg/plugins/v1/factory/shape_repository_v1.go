package factory

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/plugins/v1/shapes"
)

const (
	Circle  core.ShapeType = "circle"
	Square                 = "square"
	Haxagon                = "haxagon"
)

type ShapeRepositoryV1 struct{}

func (repo *ShapeRepositoryV1) Shapes() map[core.ShapeType]core.Shape {
	return map[core.ShapeType]core.Shape{
		Circle:  shapes.Circle{},
		Square:  shapes.Square{},
		Haxagon: shapes.Hexagon{},
	}
}
