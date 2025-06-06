package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("salvando os dados do usuario %s no banco\n", u.nome)

}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"iago", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"davi", 40}
	fmt.Println(usuario2)
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)
}
