package grokking

import (
	"golang.org/x/exp/constraints"
)

func BinarySearch[K constraints.Ordered](list []K, target K) (int, bool) {
	len := len(list)

	lo, hi := 0, len
	for lo < hi {
		mid := int(uint(lo+hi) >> 1) //avoid overflow

		if list[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo, lo < len && list[lo] == target
}
