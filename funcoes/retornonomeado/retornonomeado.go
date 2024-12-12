package main

import "fmt"

func trocar(p1, p2 string) (segundo, primeiro string) {
	primeiro = p1
	segundo = p2
	return // retorno limpo
}

func main() {
	x, y := trocar("Primeiro", "Segundo")
	fmt.Println(x, y)

	segundo, primeiro := trocar("Primeiro", "Segundo")
	fmt.Println(primeiro, segundo)
}
