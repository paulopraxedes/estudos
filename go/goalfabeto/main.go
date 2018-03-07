package main

import (
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		goalfabeto.MostrarTabela()
		return
	} else if len(os.Args) == 2 {
		goalfabeto.MontaAlfabeto(strings.ToLower(string(os.Args[1])), "")
		return
	} else {
		goalfabeto.MontaAlfabeto(strings.ToLower(string(os.Args[1])), strings.ToUpper(string(os.Args[2])))
		return
	}
}
