package main

import (
	"fmt"
)

func main() {
	var A, B, C float32

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	media := ((A * 2) + (B * 3) + (C * 5)) / (10)

	s := fmt.Sprintf("%.1f", media)
	fmt.Println("MEDIA =", s)
}
