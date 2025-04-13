package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtraçao int) {
	soma = n1 + n2
	subtraçao = n1 - n2
	return
}

func main() {
	soma, subtraçao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtraçao)

}
