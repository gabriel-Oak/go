package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{4.0, 4.0}

	fmt.Println(Distancia(p1, p2))
}
