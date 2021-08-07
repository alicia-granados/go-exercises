package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	areaTriangulo := (A * C) / 2
	areaCirculo := 3.14159 * (math.Pow(C, 2))
	areaTrapecio := ((A + B) * C) / 2
	areaCuadrado := math.Pow(B, 2)
	areaRectangulo := A * B

	s := fmt.Sprintf("%.3f", areaTriangulo)
	fmt.Println("TRIANGULO:", s)

	s1 := fmt.Sprintf("%.3f", areaCirculo)
	fmt.Println("CIRCULO:", s1)

	s2 := fmt.Sprintf("%.3f", areaTrapecio)
	fmt.Println("TRAPEZIO:", s2)

	s3 := fmt.Sprintf("%.3f", areaCuadrado)
	fmt.Println("QUADRADO:", s3)

	s4 := fmt.Sprintf("%.3f", areaRectangulo)
	fmt.Println("RETANGULO:", s4)
}
