package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	if (b > c) && (d > a) && (c+d > a+b) && c > 0 && c > 0 && (a%2 == 0) {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}
