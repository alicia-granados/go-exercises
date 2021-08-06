package main

import (
	"fmt"
)

func main() {
	var A string
	var B, C float32

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	salario := B + C*0.15
	s := fmt.Sprintf("%.2f", salario)
	fmt.Println("TOTAL = R$", s)
}
