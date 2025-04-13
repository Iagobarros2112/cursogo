package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "ola mundo"
	canal <- "progamando em go"

	mensagem := <-canal
	fmt.Println(mensagem)
}
