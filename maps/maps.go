package main

import "fmt"

func main() {
	fmt.Println("maps")

	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario)
	// fmt.Println(usuario["nome"]) caso eu queria selecionar somente um item

	//maps dentro de maps

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":   "engenharia mecanica",
			"campus": "campus 12",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome") // deletar uma chave do map usuario2
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}

	fmt.Println(usuario2)
}
