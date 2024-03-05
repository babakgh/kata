package core_test

import (
	"fmt"
	"log"
	// "reflect"
	"testing"

	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
)

type ShapeFactoryMock struct {
	*core.BaseShapeFactory
}

type CircleMock struct{}

func (c CircleMock) Draw() {
}

type SquareMock struct{}

func (s SquareMock) Draw() {
}

const (
	Circle core.ShapeType = "circle"
	Square                = "square"
)

type ShapeRepositoryMock struct{}

func (repo *ShapeRepositoryMock) Shapes() map[core.ShapeType]core.Shape {
	return map[core.ShapeType]core.Shape{
		Circle: CircleMock{},
		Square: SquareMock{},
	}
}

func TestAllShapeNamesShouldBeShape(t *testing.T) {
	// Setup
	log.Printf("When given a shape name, it should be a Shape")
	// Arrange
	var factory = &ShapeFactoryMock{BaseShapeFactory: &core.BaseShapeFactory{Repo: &ShapeRepositoryMock{}}}
	shapes := factory.GetShapeTypes()
	for _, shapeType := range shapes {
		t.Run(fmt.Sprintf("Test%s", strcase.ToCamel(string(shapeType))), func(t *testing.T) {
			// Act
			shape, err := factory.MakeShape(shapeType)
			if err != nil {
				t.Errorf("incorrect, got: %v, want nil", err)
				return
			}
			assert.Implements(t, (*core.Shape)(nil), shape)
		})
	}
}

func TestMakeInvalidShape(t *testing.T) {
	// Setup
	log.Printf("When app creats a blah shape, it returns error")
	// Arrange
	want := core.ErrShapeIsNotFound
	var factory = &ShapeFactoryMock{BaseShapeFactory: &core.BaseShapeFactory{Repo: &ShapeRepositoryMock{}}}
	// Act
	_, got := factory.MakeShape("blah")
	// Assert
	assert.Equal(t, want, got)
}
