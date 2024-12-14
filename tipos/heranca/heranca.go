package main

import "fmt"

type Animal struct {
	Nome string
}

func (a Animal) Falar() {
	fmt.Println("Som...")
}

type Cachorro struct {
	Animal // Campos anonimos
	peso   float64
}

func (c Cachorro) Falar() {
	fmt.Println("Au au!")
}

func main() {
	animal := Animal{"Animal"}
	animal.Falar()

	cachorro := Cachorro{Animal{"Cachorro"}, 10}
	cachorro.Falar()
	cachorro.Animal.Falar()
	fmt.Println(cachorro)
}
