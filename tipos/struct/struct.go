package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	type pessoa struct {
		nome  string
		idade int
	}

	var p pessoa
	p.nome = "João"
	p.idade = 20

	fmt.Println(p.nome)
	fmt.Println(p.idade)

	produto := produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto)
	fmt.Println(produto.precoComDesconto())
}
