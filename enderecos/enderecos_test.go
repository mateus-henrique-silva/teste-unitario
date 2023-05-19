package enderecos

import (
	"testing"
)

type enderecos struct {
	enderecosInseridos string
	endercosEsperados  string
}

func TestEnderecoFunc(t *testing.T) {
	cenarioDoTest := []enderecos{
		{"Avenid Brasil", "Avenida"}, {"Avenida Paulista", "Avenida"},
	}
	for _, cenario := range cenarioDoTest {
		retornoDeValor := TipoDeEnderecos(cenario.endercosEsperados)
		if retornoDeValor != cenario.endercosEsperados {

			t.Errorf("O tipo recebido é diferente do esperado que é %s e recebeu %s", cenario.endercosEsperados, retornoDeValor)
		}
	}

}
