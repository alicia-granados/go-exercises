package main

import (
	"fmt"
	"sort"
)

func main() {

	var A, B, C int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	x := []int{A, B, C}

	sort.Ints(x)

	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
