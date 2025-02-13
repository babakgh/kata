package main

import "fmt"

// Given an array of integers, return the indices of the two numbers that add up to a given target.
// There is always exactly one solution
// Input: nums = [5, 2, 3], target = 8 Output: [0, 2]

func main() {
	list := []int{4, 5, 2, 3}
	target := 8

	var a [2]int
	h := map[int]int{}
	for i, num := range list {
		com := target - num
		if j, ok := h[com]; ok {
			a = [2]int{j, i}
			break
		} else {
			h[num] = i
		}
	}

	fmt.Println(a)
}
