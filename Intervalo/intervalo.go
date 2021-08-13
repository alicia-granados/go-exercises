package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Scanln(&a)

	if a < 0 || a > 100 {
		fmt.Println("Fora de intervalo")
	} else if a == 0 || a <= 25.0000 {
		fmt.Println("Intervalo [0,25]")
	} else if a > 25.0000 && a <= 50.0000 {
		fmt.Println("Intervalo (25,50]")
	} else if a > 75 && a <= 100.000 {
		fmt.Println("Intervalo (75,100]")
	}
}
