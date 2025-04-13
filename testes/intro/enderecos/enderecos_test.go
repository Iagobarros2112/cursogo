//teste de unidade

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereço(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Praça Das Rosas", "tipo invalido"},
		{"Estrada Qualqer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "avenida"},
		{" ", "tipo invalido"},
		
	}

	for_, cenario := range cenarioscenariosDeTeste {
		tipoDeTipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
            t.Errorf("o tipo recebido %s é diferente do esperado %s", 
			    tipoDeEnderecoRecebido,
		        cenario.retornoEsperado,

		    )
		}
	}


}
