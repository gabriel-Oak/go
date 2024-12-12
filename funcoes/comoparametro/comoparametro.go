package main

func multiplicacao(a, b int) int {
	return a * b
}

func exec(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	println(resultado)
}
