package grokking_test

import (
	"fmt"
	"log"
	"math"
	"testing"

	. "github.com/babakgh/kata/algorithms/go/pkg/grokking"
)

func TestBinarySearch_int(t *testing.T) {
	log.Printf("When search on int numbers, it finds the number in the list")
	// Arrange
	list := createListInt(100)

	binarySearches := []struct {
		name string
		f    func(list []int, item int) int
	}{
		{"iteration", BinarySearchByIteration[int]},
		{"recursion", BinarySearchByRecursion[int]},
	}

	for _, binarySearch := range binarySearches {
		t.Run(fmt.Sprintf("BinarySearchBy%s", binarySearch.name), func(t *testing.T) {
			// Act-Assert
			assertEqual(t, binarySearch.f(nil, 1), -1)
			assertEqual(t, binarySearch.f([]int{1}, 1), 0)
			assertEqual(t, binarySearch.f(list, 50), 49)
			assertEqual(t, binarySearch.f(list, 75), 74)
			assertEqual(t, binarySearch.f(list, 25), 24)
			assertEqual(t, binarySearch.f(list, 72), 71)
			assertEqual(t, binarySearch.f(list, 100), 99)
			assertEqual(t, binarySearch.f(list, 1), 0)
			assertEqual(t, binarySearch.f(list, 13), 12)
			assertEqual(t, binarySearch.f(list, 5), 4)
			assertEqual(t, binarySearch.f(list, 101), -1)
		})
	}
}

func TestBinarySearch_float(t *testing.T) {
	log.Printf("When search on float numbers, it finds the float in the list")
	// Arrange
	list := createListFloat32(100)

	binarySearches := []struct {
		name string
		f    func(list []float32, item float32) int
	}{
		{"iteration", BinarySearchByIteration[float32]},
		{"recursion", BinarySearchByRecursion[float32]},
	}

	for _, binarySearch := range binarySearches {
		t.Run(fmt.Sprintf("BinarySearchBy%s", binarySearch.name), func(t *testing.T) {
			// Act-Assert
			assertEqual(t, binarySearch.f(nil, 1), -1)
			assertEqual(t, binarySearch.f([]float32{1}, 1), 0)
			assertEqual(t, binarySearch.f(list, 50.5), 49)
			assertEqual(t, binarySearch.f(list, 75.5), 74)
			assertEqual(t, binarySearch.f(list, 25.5), 24)
			assertEqual(t, binarySearch.f(list, 72.5), 71)
			assertEqual(t, binarySearch.f(list, 100.5), 99)
			assertEqual(t, binarySearch.f(list, 1.5), 0)
			assertEqual(t, binarySearch.f(list, 13.5), 12)
			assertEqual(t, binarySearch.f(list, 5.5), 4)
			assertEqual(t, binarySearch.f(list, 101.5), -1)
			assertEqual(t, binarySearch.f(list, 50.1), -1)
		})
	}
}

func TestBinarySearchByRecursion_string(t *testing.T) {
	log.Printf("When search on strings, it finds the string in the list")
	// Arrange
	list := make([]string, 26)
	for i := range list {
		list[i] = string(rune('a' + i))
	}

	binarySearches := []struct {
		name string
		f    func(list []string, item string) int
	}{
		{"iteration", BinarySearchByIteration[string]},
		{"recursion", BinarySearchByRecursion[string]},
	}

	for _, binarySearch := range binarySearches {
		t.Run(fmt.Sprintf("BinarySearchBy%s", binarySearch.name), func(t *testing.T) {
			// Act-Assert
			assertEqual(t, binarySearch.f(nil, "a"), -1)
			assertEqual(t, binarySearch.f([]string{"a"}, "a"), 0)
			assertEqual(t, binarySearch.f(list, "m"), 12)
			assertEqual(t, binarySearch.f(list, "r"), 17)
			assertEqual(t, binarySearch.f(list, "g"), 6)
			assertEqual(t, binarySearch.f(list, "i"), 8)
			assertEqual(t, binarySearch.f(list, "z"), 25)
			assertEqual(t, binarySearch.f(list, "w"), 22)
			assertEqual(t, binarySearch.f(list, "aa"), -1)
			assertEqual(t, binarySearch.f(list, "za"), -1)
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	binarySearches := []struct {
		name string
		f    func(list []int, item int) int
	}{
		{"iteration", BinarySearchByIteration[int]},
		{"recursion", BinarySearchByRecursion[int]},
	}

	for _, binarySearch := range binarySearches {
		for k := 0.; k <= 10; k++ {
			n := int(math.Pow(2, k))
			list := createListInt(n)
			b.Run(fmt.Sprintf("%s/%d", binarySearch.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					binarySearch.f(list, 0)
					binarySearch.f(list, n-1)
					binarySearch.f(list, n/2)
					binarySearch.f(list, n/4+1)
				}
			})
		}
	}
}

func assertEqual[K comparable](t *testing.T, got, want K) {
	t.Helper()
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
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
