package main

import (
	"fmt"
)

func main() {
	var A, B float32

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	media := ((A * 3.5) + (B * 7.5)) / (3.5 + 7.5)

	s := fmt.Sprintf("%.5f", media)
	fmt.Println("MEDIA = ", s)
}
