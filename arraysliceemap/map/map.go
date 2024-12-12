package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98324324255] = "Pedro"
	aprovados[32412312456] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[32412312456])
	delete(aprovados, 32412312456)
	fmt.Println(aprovados[32412312456])
}
