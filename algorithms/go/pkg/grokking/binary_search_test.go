package grokking_test

import (
	"testing"

	. "github.com/babakgh/kata/algorithms/go/pkg/grokking"
)

func TestBinarySearch(t *testing.T) {
	t.Logf("When value is procided, it finds the orderd poistion") //TODO better caption

	list := createListInt(100)

	var tests = []struct {
		list   []int
		target int
		pos    int
		found  bool
	}{
		{[]int{}, 1, 0, false},
		{[]int{1}, 2, 1, false},
		{[]int{1, 3}, 2, 1, false},
		{[]int{1, 1}, 2, 2, false},
		{[]int{2, 2}, 1, 0, false},
		{[]int{1, 1, 2, 2, 4, 5}, 3, 4, false},
		{[]int{1, 2, 2, 2}, 2, 1, true},
		//
		{nil, 1, 0, false},
		{[]int{1}, 1, 0, true},
		{[]int{1}, 5, 1, false},
		{[]int{1, 2}, 1, 0, true},
		{[]int{1, 2, 2, 2, 2, 3, 4, 5, 6}, 2, 1, true},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 1, 0, true},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 2, 7, false},
		{[]int{1, 2}, 2, 1, true},
		{[]int{-2, -1, 0, 1, 2}, -1, 1, true},
		{list, 50, 49, true},
		{list, 75, 74, true},
		{list, 25, 24, true},
		{list, 72, 71, true},
		{list, 100, 99, true},
		{list, 1, 0, true},
		{list, 13, 12, true},
		{list, 5, 4, true},
		{list, 101, 100, false},
	}

	for _, e := range tests {
		pos, found := BinarySearch(e.list, e.target)

		got := []any{pos, found}
		want := []any{e.pos, e.found}
		assertDeepEqual(t, got, want)
	}
}

func createListInt(n int) []int {
	list := make([]int, n)
	for i := range list {
		list[i] = i + 1
	}
	return list
}

func createListFloat32(n int) []float32 {
	list := make([]float32, n)
	for i := range list {
		list[i] = float32(i) + 1 + 0.5
	}
	return list
}

func createListString() []string {
	list := make([]string, 26)
	for i := range list {
		list[i] = string(rune('a' + i))
	}

	return list
}
