package core

import (
	"errors"
)

var Shapes = map[string]interface{}{
	"circle":  Circle{},
	"square1": Square{},
}

type ShapeFactoryImp struct{}

var ErrShapIsNotFound = errors.New("ShapeIsNotFound")

func (shi *ShapeFactoryImp) Make(shapeName string) (interface{}, error) {
	shape, ok := Shapes[shapeName]
	if !ok {
		return nil, ErrShapIsNotFound
	}
	return shape, nil
}

func (shi *ShapeFactoryImp) GetShapeNames() []string {
	shapeNames := make([]string, 0, len(Shapes))
	for name := range Shapes {
		shapeNames = append(shapeNames, name)
	}
	return shapeNames
}
