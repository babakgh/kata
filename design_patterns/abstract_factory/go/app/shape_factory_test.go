package app_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/babakgh/kata/design_patterns/abstract_factory/go/app"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

func TestMakeCircle(t *testing.T) {
	// Setup
	log.Printf("When app creats a circle shape, it is a core.Circle")
	// Arrange
	factory := &core.ShapeFactoryImp{}
	circle, _ := factory.Make("circle")
	c := circle.(app.Shape)
	c.Draw()
	want := reflect.TypeFor[core.Circle]()
	// Act
	got := reflect.TypeOf(circle)
	// Assert
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}
