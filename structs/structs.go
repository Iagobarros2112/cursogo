package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereço endereço
}

type endereço struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("arquivo structs")

	var u usuario
	u.nome = "davi"
	u.idade = 21
	fmt.Println(u)

	endereçoexemplo := endereço{"rua das acacias", 0}

	usuario2 := usuario{
		"davi", 21, endereçoexemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "davi"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 22}
	fmt.Println(usuario4)
}
