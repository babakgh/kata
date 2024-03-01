package main

import "github.com/babakgh/kata/design_patterns/abstract_factory/go/app"
import "github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"

func main() {
	var shapeFactory = &core.ShapeFactoryImp{}

	app := app.App{}
	app.Start(shapeFactory)
}
