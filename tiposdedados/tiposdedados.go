package main

import "fmt"

func main() {

	//int, uint, int8, int16, int32, int64
	var numero int64 = -10000000 //int64
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2) //insigned int

	//alias
	//var numero3 int32 =12456
	var numero3 rune = 12456
	fmt.Println(numero3)

	//byte =uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)

	var numeroreal2 float64 = 1230000000.45
	fmt.Println(numeroreal2)

	numeroreal3 := 12345.67
	fmt.Println(numeroreal3)

	//strings

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char)

	//valores vazios

	var texto int16
	fmt.Println(texto)

	var texto2 string
	fmt.Println(texto2)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
}
