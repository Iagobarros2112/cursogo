package main

import (
	"fmt"
	//"time"
)

func main() {
	//	i := 0

	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//	}

	//	fmt.Println(i)

	//for j := 0; j < 10; j++ {
	//	fmt.Println("incrementando j", j)
	//	time.Sleep(time.Second)

	//}

	nomes := [3]string{"joao", "davi", "lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
