package main

import (
	"fmt"

	"github.com/gabriel-Oak/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - mistura mensagens de dois canais em um terceiro canal
func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	saida := make(chan string)
	go encaminhar(entrada1, saida)
	go encaminhar(entrada2, saida)
	return saida
}

func main() {
	c := multiplexar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c + " | " + <-c)
	fmt.Println(<-c + " | " + <-c)
}
