package grokking_test

import (
	// "log"
	// "reflect"
	"testing"

	. "github.com/babakgh/kata/algorithms/go/pkg/grokking"
)

func TestDijkstra(t *testing.T) {
	// Setup
	t.Logf("When search on shortest weigthed path, it finds it")
	// Arrange
	graph := map[string]map[string]int{}

	graph["a"] = map[string]int{"b": 4, "c": 10}
	graph["b"] = map[string]int{"d": 21}
	graph["c"] = map[string]int{"e": 5, "f": 12}
	graph["d"] = map[string]int{"g": 4}
	graph["e"] = map[string]int{"d": 5}
	graph["f"] = map[string]int{"d": 12}

	// Act-Assert
	assertDeepEqual(t, Dijkstra(graph, "z", "f"), []string{})
	assertDeepEqual(t, Dijkstra(graph, "a", "g"), []string{"a", "c", "e", "d", "g"})
}
