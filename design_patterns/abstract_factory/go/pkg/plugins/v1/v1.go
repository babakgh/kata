package main

import "github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
import "github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/plugins/v1/factory"

// exported as symbol
var Factory factory.ShapeFactoryV1 = factory.ShapeFactoryV1{BaseShapeFactory: &core.BaseShapeFactory{Repo: &factory.ShapeRepositoryV1{}}}
