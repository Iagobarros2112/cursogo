package main

import (
	"comando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("inicio")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
