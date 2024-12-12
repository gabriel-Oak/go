package main

import (
	"fmt"
	"runtime/debug"
)

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
	fmt.Println("Main")
}
