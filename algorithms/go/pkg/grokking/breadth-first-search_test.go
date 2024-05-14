package grokking_test

import (
	"log"
	"reflect"
	"testing"

	. "github.com/babakgh/kata/algorithms/go/pkg/grokking"
)

func TestBreadthFirstSearch(t *testing.T) {
	// Setup
	log.Printf("When search on shortest path, it finds it")
	// Arrange
	graph := map[string][]string{}

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["alice"] = []string{"peggy", "you"}
	graph["bob"] = []string{"anuj", "peggy", "you"}
	graph["claire"] = []string{"thom", "jonny", "you"}
	graph["peggy"] = []string{"alice", "bob"}
	graph["anuj"] = []string{"bob"}
	graph["thom"] = []string{"claire"}
	graph["jonny"] = []string{"ella", "claire"}

	// Act-Assert
	assertDeepEqual(t, BFS(graph, "unknown", "jonny"), []string{})
	assertDeepEqual(t, BFS(graph, "unknown", "unknown"), []string{})
	assertDeepEqual(t, BFS(graph, "you", "unknown"), []string{})
	assertDeepEqual(t, BFS(graph, "you", "you"), []string{"you"})
	assertDeepEqual(t, BFS(graph, "you", "jonny"), []string{"you", "claire", "jonny"})
	assertDeepEqual(t, BFS(graph, "you", "alice"), []string{"you", "alice"})
	assertDeepEqual(t, BFS(graph, "you", "bob"), []string{"you", "bob"})
	assertDeepEqual(t, BFS(graph, "you", "peggy"), []string{"you", "alice", "peggy"})
	assertDeepEqual(t, BFS(graph, "you", "anuj"), []string{"you", "bob", "anuj"})
	assertDeepEqual(t, BFS(graph, "you", "ella"), []string{"you", "claire", "jonny", "ella"})
}

func assertDeepEqual[T any](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
