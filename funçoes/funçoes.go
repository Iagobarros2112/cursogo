package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtraçao := n1 - n2
	return soma, subtraçao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da funçao 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtraçao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtraçao)
}
