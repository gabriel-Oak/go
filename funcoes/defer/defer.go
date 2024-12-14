package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	defer fmt.Println("1.0")
	defer fmt.Println("1.1")
	defer fmt.Println("1.2")
	defer fmt.Println("1.3")

	fmt.Println("2")
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
