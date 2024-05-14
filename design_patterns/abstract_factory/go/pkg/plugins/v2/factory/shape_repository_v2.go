package factory

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/plugins/v2/shapes"
)

const (
	Circle              core.ShapeType = "circle"
	EquilateralTriangle                = "equilateral_triangle"
)

var allShapes = map[core.ShapeType]core.Shape{
	Circle:              shapes.Circle{},
	EquilateralTriangle: shapes.EquilateralTriangle{},
}

type ShapeRepositoryV2 struct{}

func (repo *ShapeRepositoryV2) Shapes() map[core.ShapeType]core.Shape {
	return allShapes
}
