package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	println(tv50, tv32, sorvete)

	tv50, tv32, sorvete = compras(true, false)
	println(tv50, tv32, sorvete)

	tv50, tv32, sorvete = compras(false, true)
	println(tv50, tv32, sorvete)

	tv50, tv32, sorvete = compras(false, false)
	println(tv50, tv32, sorvete)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t\n", tv50, tv32, sorvete)
}
