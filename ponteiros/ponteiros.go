package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var variavel1 int = 10

	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++

	//ponteiros sao referencias de memoria
	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro) //desreferenciacao

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) //desrefenrenciacao
}
