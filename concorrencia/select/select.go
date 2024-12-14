package main

import (
	"time"

	"github.com/gabriel-Oak/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// Estrutura de controle para concorrência
	// A estrutura select é utilizada para concorrência
	// O select é uma estrutura de controle que permite a execução
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Millisecond * 900):
		return "Todos perderam!"
		// Caso não tenha resposta em nenhum dos canais
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
	)
	println(campeao)
}
