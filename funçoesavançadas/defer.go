package main

import "fmt"

func funcao1() {
	fmt.Println("exe funçao 1")

}

func funcao2() {
	fmt.Println("exe funçao 2")

}

func alunoEstaAprovado(media1, media2 float32) bool {
	fmt.Println("entrando ma funcao para verificar se o aluno esta aprovado")

	media := (media1 + media2) / 2

	if media >= 6 {
		defer fmt.Println("media calculada o resltado sera mostrado na tela")
		return true
	}

	return false
}

func main() {

	fmt.Println(alunoEstaAprovado(7, 8))

}
