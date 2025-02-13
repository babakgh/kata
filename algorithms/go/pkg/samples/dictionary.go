package main

import (
	"fmt"

	"github.com/babakgh/kata/algorithms/go/pkg/grokking"
)

func main() {

	list := []string{"Babak", "Nima", "Shadan"}

	fmt.Println(grokking.BinarySearch(list, "Ali"))
	fmt.Println(grokking.BinarySearch(list, "Babak"))
	fmt.Println(grokking.BinarySearch(list, "Monir"))
	fmt.Println(grokking.BinarySearch(list, "Salva"))

}
