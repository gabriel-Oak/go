package main

import "fmt"

func rotina(c chan int) {
	fmt.Println("Executou")
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
}

func main() {
	c := make(chan int, 3)
	go rotina(c)

	fmt.Println(<-c)
}
