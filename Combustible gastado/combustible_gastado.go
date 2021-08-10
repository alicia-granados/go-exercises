package main

import (
	"fmt"
)

func main() {
	var tiempoDuracion, velocidadMedia int

	fmt.Scanln(&tiempoDuracion)
	fmt.Scanln(&velocidadMedia)

	fmt.Println(combustibleGastado(tiempoDuracion, velocidadMedia))
}
func combustibleGastado(tiempoDuracion, velocidadMedia int) string {
	distancia := float64(tiempoDuracion * velocidadMedia)
	litrosCombustibleGastado := distancia / 12
	s := fmt.Sprintf("%.3f", litrosCombustibleGastado)
	return fmt.Sprint(s)
}
