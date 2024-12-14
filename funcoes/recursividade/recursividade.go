package main

import "fmt"

func fatorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func fatorial2(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial2(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	fmt.Println(fatorial(5)) // 120
	resultado, _ := fatorial2(5)
	fmt.Println(resultado) // 120
	_, err := fatorial2(-4)
	if err != nil {
		fmt.Println(err)
	}
}
