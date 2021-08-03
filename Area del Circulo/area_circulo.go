package main

import (
	"fmt"
	"math"
)

func main() {
	var radio, area, pi float64
	fmt.Scanln(&radio)
	pi = 3.14159
	area = (pi) * (math.Pow(radio, 2))
	fmt.Println(fmt.Sprintf("A=%.4f", area))
}
