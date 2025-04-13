package main

import (
	"fmt"

	"intro/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Pailista")
	fmt.Println(tipoEndereco)
}
