package main

import (
	"fmt"
)

func main() {
	var dinero, aux int

	fmt.Scanln(&dinero)

	billeteCien := dinero / 100
	aux = dinero % 100
	billeteCincuenta := aux / 50
	aux = aux % 50
	billeteVeinte := aux / 20
	aux = aux % 20
	billeteDiez := aux / 10
	aux = aux % 10
	billeteCinco := aux / 5
	aux = aux % 5
	billeteDos := aux / 2
	aux = aux % 2
	billeteUno := aux / 1
	aux = aux % 1

	fmt.Println(dinero)
	fmt.Println(billeteCien, "nota(s) de R$ 100,00")
	fmt.Println(billeteCincuenta, "nota(s) de R$ 50,00")
	fmt.Println(billeteVeinte, "nota(s) de R$ 20,00")
	fmt.Println(billeteDiez, "nota(s) de R$ 10,00")
	fmt.Println(billeteCinco, "nota(s) de R$ 5,00")
	fmt.Println(billeteDos, "nota(s) de R$ 2,00")
	fmt.Println(billeteUno, "nota(s) de R$ 1,00")

}
