package main

import (
	"fmt"
)

func main() {
	var A int
	var B float32

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	fmt.Println(consumo(A, B))
}
func consumo(a int, b float32) string {
	consumoPromedio := float32(a) / b
	s := fmt.Sprintf("%.3f", consumoPromedio)
	return fmt.Sprint(s, " km/l")
}
