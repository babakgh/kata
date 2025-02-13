package main

import (
	"fmt"
	"sort"
	"strings"
)

// findStart finds the first index where words start with the prefix
func findStart(words []string, prefix string) int {
	var cmp = func(i int) int {
		return strings.Compare(prefix, words[i])
	}

	i, _ := sort.Find(len(words), cmp)
	return i
	// return sort.Search(len(words), func(i int) bool {
	// 	fmt.Println(words[i])
	// 	return strings.HasPrefix(words[i], prefix) || words[i] >= prefix
	// })
}

// findEnd finds the first index where words no longer match the prefix
func findEnd(words []string, prefix string) int {
	var cmp = func(i int) int {
		fmt.Printf("%v %v\n", prefix, words[i])
		if !strings.HasPrefix(words[i], prefix) {
			return 0
		} else {
			return strings.Compare(words[i], prefix)
		}
	}

	i, _ := sort.Find(len(words), cmp)
	return i

	// return sort.Search(len(words), func(i int) bool {
	// 	fmt.Println(words[i])
	// return !strings.HasPrefix(words[i], prefix)
	// })
}

// autocomplete returns all words matching the prefix
func autocomplete(words []string, prefix string) []string {
	// Binary search for the range of words matching the prefix
	start := findStart(words, prefix)
	end := findEnd(words, prefix)

	fmt.Printf("%v:%v\n", start, end)
	// Return the matching words
	return words[start:end]
}

func main() {
	// Sorted list of words
	words := []string{"apple", "apply", "banana", "band", "bandage", "bandana", "cat", "catalog", "dog"}

	fmt.Printf("Autocomplete results for '%s': %v\n", "app", autocomplete(words, "app"))
	fmt.Printf("Autocomplete results for '%s': %v\n", "ban", autocomplete(words, "ban"))

}
