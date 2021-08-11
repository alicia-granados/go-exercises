package main

import (
	"fmt"
)

func main() {
	var edadDias int

	fmt.Scanln(&edadDias)

	ano := edadDias / 365
	edadDias = edadDias - (ano * 365)

	mes := edadDias / 30
	edadDias = edadDias - (mes * 30)

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", ano, mes, edadDias)

}
