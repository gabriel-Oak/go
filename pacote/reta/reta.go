package main

import (
	"math"
)

// Inicializando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)!
// Inicializando com letra minúscula é PRIVADO (visível apenas dentro do pacote)!

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (catX float64, catY float64) {
	catX = math.Abs(a.x - b.x)
	catY = math.Abs(a.y - b.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	catX, catY := catetos(a, b)
	return math.Sqrt(math.Pow(catX, 2) + math.Pow(catY, 2))
}

func calcDistancia(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}
