package enderecos

import "strings"

func TipoDeEnderecos(enderecos string) string {
	tiposdeend := []string{"rua", "avenida", "estradas", "rodovia"}
	endMaiuscula := strings.ToLower(enderecos)
	primeirapalavra := strings.Split(endMaiuscula, " ")[0]
	endtype := false
	for _, tipo := range tiposdeend {

		if tipo == primeirapalavra {
			endtype = true
		}

	}
	if endtype {
		return strings.Title(primeirapalavra)
	}

	return "tipo invalido"
}
