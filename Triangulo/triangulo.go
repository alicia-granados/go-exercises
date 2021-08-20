package main

import (
	"fmt"
)

func main() {

	var A, B, C float64

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	if (A+B > C) && (B+C > A) && (A+C > B) {
		perimetroTriangulo := A + B + C
		s2 := fmt.Sprintf("%.1f", perimetroTriangulo)
		fmt.Println("Perimetro =", s2)
	} else {
		areaTrapecio := ((A + B) * C) / 2
		s2 := fmt.Sprintf("%.1f", areaTrapecio)
		fmt.Println("Area =", s2)
	}

}
