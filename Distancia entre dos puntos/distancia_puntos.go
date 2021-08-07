package main

import (
	"fmt"
	"math"
)

func main() {
	var X1, Y1, X2, Y2 float64

	fmt.Scan(&X1)
	fmt.Scan(&Y1)

	fmt.Scan(&X2)
	fmt.Scan(&Y2)

	x := (math.Pow((X2 - X1), 2))
	y := (math.Pow((Y2 - Y1), 2))

	distanciaPunto := math.Sqrt(x + y)
	s := fmt.Sprintf("%.4f", distanciaPunto)
	fmt.Println(s)
}
