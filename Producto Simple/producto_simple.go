package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	X := A * B
	fmt.Println("PROD =", X)
}
