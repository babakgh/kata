package core

import "errors"

var ErrShapeIsNotFound = errors.New("ShapeIsNotFound")

type ShapeFactory interface {
	MakeShape(shapeType ShapeType) (Shape, error)
	GetShapeTypes() []ShapeType
}
