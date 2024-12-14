package main

import (
	"fmt"
	"time"
)

func isPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Primos - retorna um canal
func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}

	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(90, c)

	for primo := range c {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")
}
