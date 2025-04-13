package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	default:
		return "numero invalido"
	}

}

func main() {
	fmt.Println("switch")

	dia := diaDaSemana(1)
	fmt.Println(dia)
}
