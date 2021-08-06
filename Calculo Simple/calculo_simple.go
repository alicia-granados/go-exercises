package main

import (
	"fmt"
)

func main() {
	var A, B int
	var C float32

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	calculo := (float32(B) * C)

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	calculo += (float32(B) * C)
	s := fmt.Sprintf("%.2f", calculo)
	fmt.Println("VALOR A PAGAR: R$", s)
}
