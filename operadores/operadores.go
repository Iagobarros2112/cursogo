package main

import "fmt"

func main() {
	//operadores aritimeticos
	soma := 1 + 2
	subtraçao := 1 - 2
	divisao := 10 / 4
	multiplicaçao := 10 * 5
	restodadivisao := 10 % 2

	fmt.Println(soma, subtraçao, divisao, multiplicaçao, restodadivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// fim dos operadores aritimeticos

	//atribuiçao
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	//fim operadores de atribuiçao

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// fim dos operadores relacionais

	//operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //and
	fmt.Println(verdadeiro || falso) //ou
	fmt.Println(!verdadeiro)         //not
	fmt.Println(!falso)              //not

	//operadores unarios
	numero := 10
	numero++
	numero += 15

	fmt.Println(numero)

	numero--
	numero -= 26

	//fim operadores unarios

	//operadores ternarios
	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor qur 5"
	}

	fmt.Println(texto)
}
