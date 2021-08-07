package main

import (
	"fmt"
)

func main() {
	var A, B, C int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	fmt.Println(numeroMayor(A, B, C))
}
func numeroMayor(a, b, c int) string {
	numMayor := a
	if a < b || a < c {
		if b < c {
			numMayor = c
		} else {
			numMayor = b
		}
	}
	return fmt.Sprint(numMayor, " eh o maior")
}
