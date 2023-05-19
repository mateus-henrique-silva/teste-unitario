package enderecos

import "testing"

func TestEnderecoFunc(t *testing.T) {
	entrada := "Avenida Paulista"
	esperado := "Avenida"

	TipoRecebido := TipoDeEnderecos(entrada)

	if TipoRecebido != esperado {

		t.Errorf("O tipo recebido é diferente do esperado que é %s e recebeu %s", esperado, TipoRecebido)
	}
}
