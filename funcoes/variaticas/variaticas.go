package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1, 5.9, 10.0))
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1))
	fmt.Printf("Média: %.2f\n", media(7.7))
	fmt.Printf("Média: %.2f\n", media())
}
