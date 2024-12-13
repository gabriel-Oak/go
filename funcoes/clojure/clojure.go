package main

import "fmt"

func clojure() func() {
	x := 10
	return func() {
		fmt.Println(x)
	}
}

func main() {
	x := 20
	fmt.Println(x) // 20

	imprimeX := clojure()
	imprimeX() // 10
}
