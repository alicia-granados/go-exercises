package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	bCuadrado := math.Pow(b, 2)
	raiz := math.Sqrt(bCuadrado - (4 * a * c))

	if (bCuadrado-(4*a*c)) < 0 || a == 0 {
		fmt.Println("Impossivel calcular")
	} else {
		fmt.Printf("R1 = %.5f\nR2 = %.5f\n", ((-b + raiz) / (2 * a)), ((-b - raiz) / (2 * a)))
	}

}
