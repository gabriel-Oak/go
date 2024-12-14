package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 900)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
				break
			case s := <-entrada2:
				c <- s
				break
			case <-time.After(time.Second):
				close(c)
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	// fmt.Println(<-c, <-c)
	// fmt.Println(<-c, <-c)
	// fmt.Println(<-c, <-c)
	for resposta := range c {
		fmt.Println(resposta)
	}
}
