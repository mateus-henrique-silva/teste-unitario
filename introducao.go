package main

import (
	"fmt"

	"go.mod/enderecos"
)

func main() {
	resultado := enderecos.TipoDeEnderecos("Avenida Paulista")
	fmt.Println(resultado)
}
