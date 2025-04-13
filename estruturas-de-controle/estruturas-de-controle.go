package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("maior que 15")

	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("numero é menor que -10")
	} else {
		fmt.Println("entre 0 e -10")
	}
}
