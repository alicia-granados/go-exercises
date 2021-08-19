package main

import (
	"fmt"
)

func main() {
	var X, Y float64

	fmt.Scan(&X)
	fmt.Scan(&Y)

	if X > 0 && Y > 0 {

		fmt.Println("Q1")

	} else if X < 0 && Y > 0 {

		fmt.Println("Q2")

	} else if X < 0 && Y < 0 {

		fmt.Println("Q3")

	} else if X > 0 && Y < 0 {

		fmt.Println("Q4")

	} else if X == 0 && Y == 0 {

		fmt.Println("Origem")

	} else if Y == 0 {
		fmt.Println("Eixo X\n")
	} else if X == 0 {
		fmt.Println("Eixo Y\n")
	}

}
