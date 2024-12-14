package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bmw7 struct {
	turboLigado bool
}

func (b *bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
	b.turboLigado = true
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	carro1 := bmw7{false}
	carro1.ligarTurbo()
	println(carro1.turboLigado)

	var carro2 esportivoLuxuoso = &bmw7{false}
	carro2.ligarTurbo()
	println(carro2)
}
