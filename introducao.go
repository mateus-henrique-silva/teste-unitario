package main

import (
	"fmt"

	"go.mod/enderecos"
)

func main() {
	resultado := enderecos.TipoDeEnderecos("Avenida brasil")
	fmt.Println(resultado)
}
