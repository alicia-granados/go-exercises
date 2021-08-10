package main

import (
	"fmt"
)

func main() {
	var A int

	fmt.Scanln(&A)

	fmt.Println(distancia(A))
}
func distancia(a int) string {
	var distancia float64
	distancia = (float64(a) / ((90 / 60.0) - (60 / 60.0)))
	s := fmt.Sprintf("%.0f", distancia)
	return fmt.Sprint(s, " minutos")
}
