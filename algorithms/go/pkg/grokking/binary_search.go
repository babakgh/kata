package grokking

import (
	"golang.org/x/exp/constraints"
)

func BinarySearchByRecursion[K constraints.Ordered](list []K, item K) int {
	size := len(list)
	if size == 0 {
		return -1
	}

	return tailBinarySearch(list, item, 0, size-1)
}

func tailBinarySearch[K constraints.Ordered](list []K, item K, lo, hi int) int {
	if hi < lo {
		return -1
	}

	mid := (hi + lo) / 2

	if list[mid] == item {
		return mid
	} else if list[mid] < item {
		return tailBinarySearch(list, item, mid+1, hi)
	} else {
		return tailBinarySearch(list, item, lo, mid-1)
	}
}

func BinarySearchByIteration[K constraints.Ordered](list []K, item K) int {
	for lo, hi := 0, len(list)-1; lo <= hi; {
		mid := (lo + hi) / 2

		if list[mid] == item {
			return mid
		} else if list[mid] < item {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
