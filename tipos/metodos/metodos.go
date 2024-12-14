package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p := pessoa{
		nome:      "João",
		sobrenome: "Silva",
	}

	fmt.Println(p.nomeCompleto())

	p.setNomeCompleto("Maria Pereira")
	fmt.Println(p.nomeCompleto())
}
