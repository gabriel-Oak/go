package main

import "fmt"

// Na linguagem Go, não existe o operador ternário.
func obterResultado(nota float64) string {
	// Exemplo de uso do operador ternário em outras linguagens:
	// var resultado = (3 > 2) ? "Maior" : "Menor"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
