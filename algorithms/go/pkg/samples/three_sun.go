package main

import (
	"fmt"
	"slices"
)

func main() {
	list := []int{-1, 0, 1, 2, -1, -4}

	slices.Sort(list)
	// sum := 0

	for i, v := range list {
		fmt.Printf("%v, %v\n", v, list[i+1:])
	}

	var hi int32 = 2147483640
	var lo int32 = 2147483640
	// Overflow occurs here
	fmt.Println("Mid:", (hi+lo)/2)
	fmt.Println("Mid:", lo+(hi-lo)/2)

	a := []int{1, 2}
	b := a
	c := &a
	a[1] = 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*c)

	a = append(a, 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*c)

	tests := []any{
		[]any{[]int{}, 1, []any{0, false}},
	}

	for _, test := range tests {
		fmt.Println(test.([]any)[1])
	}
}
