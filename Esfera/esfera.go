package main

import (
	"fmt"
	"math"
)

func main() {
	var A int
	var pi float64
	pi = 3.14159

	fmt.Scan(&A)

	radio := math.Pow(float64(A), 3)
	volumen := ((1.333333333333333) * pi * radio)
	s := fmt.Sprintf("%.3f", volumen)
	fmt.Println("VOLUME =", s)
}
