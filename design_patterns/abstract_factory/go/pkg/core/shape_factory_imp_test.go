package core_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/babakgh/kata/design_patterns/abstract_factory/go/app"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
	"github.com/iancoleman/strcase"
)

func TestAllShapeNamesShouldBeShape(t *testing.T) {
	// Setup
	log.Printf("When given a shape name, it should be a Shape")
	// Arrange
	var factory = &core.ShapeFactoryImp{}
	shapes := factory.GetShapeNames()
	for _, shape_name := range shapes {
		t.Run(fmt.Sprintf("Test%s", strcase.ToCamel(shape_name)), func(t *testing.T) {
			// Act
			shape, err := factory.Make(shape_name)
			if err != nil {
				t.Errorf("incorrect, got: %v, want nil", err)
				return
			}
			want := reflect.TypeFor[app.Shape]()
			// Assert
			got := reflect.TypeOf(shape)
			if !got.Implements(want) {
				t.Errorf("incorrect, got: %v, want %v", got, want)
			}
		})
	}
}

func TestMakeInvalidShape(t *testing.T) {
	// Setup
	log.Printf("When app creats a blah shape, it returns error")
	// Arrange
	want := core.ErrShapIsNotFound
	// Act
	_, got := (&core.ShapeFactoryImp{}).Make("blah")
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
