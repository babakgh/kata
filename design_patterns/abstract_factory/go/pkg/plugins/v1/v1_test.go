package main_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/plugins/v1"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
)

func TestAllShapeNamesShouldBeShape(t *testing.T) {
	// Setup
	log.Printf("When given a shape name, it should be a Shape")
	// Arrange
	var factory = main.Factory
	shapes := factory.GetShapeTypes()
	for _, shapeType := range shapes {
		t.Run(fmt.Sprintf("Test%s", strcase.ToCamel(string(shapeType))), func(t *testing.T) {
			// Act
			shape, err := factory.MakeShape(shapeType)
			if err != nil {
				t.Errorf("incorrect, got: %v, want nil", err)
				return
			}
			// Assert
			assert.Implements(t, (*core.Shape)(nil), shape)
		})
	}
}
