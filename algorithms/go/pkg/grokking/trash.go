package grokking

import (
	"golang.org/x/exp/constraints"
)

// BinarySearch By Recursion
func BinarySearchByRecursion[K constraints.Ordered](list []K, item K) int {
	size := len(list)
	if size == 0 {
		return -1
	}

	return tailBinarySearch(list, item, 0, size-1)
}

func tailBinarySearch[K constraints.Ordered](list []K, search K, lo, hi int) int {
	if hi < lo {
		return -1
	}

	mid := lo + (hi-lo)/2 // avoid overflow

	if list[mid] == search {
		return mid
	} else if list[mid] < search {
		return tailBinarySearch(list, search, mid+1, hi)
	} else {
		return tailBinarySearch(list, search, lo, mid-1)
	}
}

// BinarySearch By Iteration
func BinarySearchByIteration[K constraints.Ordered](list []K, search K) int {
	res := -1
	for lo, hi := 0, len(list)-1; lo <= hi; {
		mid := lo + (hi-lo)/2 // avoid overflow

		if list[mid] == search {
			// return mid
			res = mid    // it is found
			hi = mid - 1 // find the first occurance
		} else if list[mid] < search {
			if res != -1 { // it is found and it is the first occurance
				break
			}
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return res
}

func InsertSorted[K constraints.Ordered](list []K, value K) []K {
	index, _ := BinarySearch(list, value)

	return append(list[:index], append([]K{value}, list[index:]...)...)
}
