package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("ola mundo", canal)

	fmt.Println("dps da funcao começar a ser executada")

	for {
		mensagem := range canal {
			fmt.Println(mensagem)
		}
	

	fmt.Println("fim do progama")

}

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
