package main

import (
	"fmt"
	"time"
)

func main() {
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("Misturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// laço infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
