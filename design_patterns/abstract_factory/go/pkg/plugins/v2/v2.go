package main

import (
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/plugins/v2/factory"
)

// exported as symbol
var Factory factory.ShapeFactoryV2 = factory.ShapeFactoryV2{BaseShapeFactory: &core.BaseShapeFactory{Repo: &factory.ShapeRepositoryV2{}}}
