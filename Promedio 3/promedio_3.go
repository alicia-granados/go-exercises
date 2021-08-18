package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, E float64

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	media := ((A * 2) + (B * 3) + (C * 4) + (D * 1)) / (10)

	s := fmt.Sprintf("%.1f", media)

	if media >= 7.0 {

		fmt.Println("Media:", s)
		fmt.Print("Aluno aprovado.\n")

	} else if media < 5.0 {

		fmt.Println("Media:", s)
		fmt.Print("Aluno reprovado.\n")

	} else if media >= 5.0 && media <= 6.9 {

		fmt.Scanln(&E)
		fmt.Println("Media:", s)
		fmt.Print("Aluno em exame.\n")
		fmt.Println("Nota do exame:", E)
		E = (media + E) / 2
		e := fmt.Sprintf("%.1f", E)

		if E >= 5.0 {
			fmt.Print("Aluno aprovado.\n")
			fmt.Println("Media final:", e)
		} else {
			fmt.Print("Aluno reprovado.\n")
			fmt.Println("Media final:", e)
		}
	}

}
