package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade string `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "dalmata ", "3"}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)

	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{
		"nome": "tobu",
		"raca": "poodle",
	}

	cachorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJson)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))

}
