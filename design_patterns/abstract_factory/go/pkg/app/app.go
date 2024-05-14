package app

import (
	"log"

	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

type app struct{}

func (a *app) Start(factory core.ShapeFactory) {
	log.Printf("Available Shapes: %v", factory.GetShapeTypes())
	for _, shape_type := range factory.GetShapeTypes() {
		if shape, err := factory.MakeShape(shape_type); err == nil {
			shape.Draw()
			log.Printf("A %T has been drawn", shape)
		} else {
			log.Print(err)
		}
	}
}

// Export app
var App app
