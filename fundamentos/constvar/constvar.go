package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// bloco de constantes
	const (
		a = 1
		b = 2
	)

	fmt.Println(a, b)

	// declarando múltiplas variáveis
	var c, d bool = true, false
	fmt.Println(c, d)

	// forma reduzida de criar múltiplas variáveis
	e, f, g := 2, false, "epa!"
	fmt.Println(e, f, g)
}
