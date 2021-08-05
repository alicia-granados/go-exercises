package main

import (
	"fmt"
)

func main() {
	var A, B int
	var C float32

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	salary := (float32(B) * C)

	fmt.Println("NUMBER =", A)

	s := fmt.Sprintf("%.2f", salary)
	fmt.Println("SALARY = U$", s)
}
