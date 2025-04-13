package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("ola mundo") //goroutine
	//go
	escrever("progamando em go")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
