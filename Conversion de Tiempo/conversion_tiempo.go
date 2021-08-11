package main

import (
	"fmt"
)

func main() {
	var segundos int

	fmt.Scanln(&segundos)

	horas := segundos / 3600
	segundos = segundos - (horas * 3600)

	minutos := segundos / 60
	segundos = segundos - (minutos * 60)

	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
