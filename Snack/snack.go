package main

import "fmt"

func main() {

	var code, amount int
	var price float64

	fmt.Scan(&code)
	fmt.Scan(&amount)

	if code == 1 {
		price = float64(amount) * 4.00
	} else if code == 2 {
		price = float64(amount) * 4.50
	} else if code == 3 {
		price = float64(amount) * 5.00
	} else if code == 4 {
		price = float64(amount) * 2.00
	} else if code == 5 {
		price = float64(amount) * 1.50
	}

	fmt.Printf("Total: R$ %.2f\n", price)

}
