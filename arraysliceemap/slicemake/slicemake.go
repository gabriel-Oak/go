package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))

	s = append(s, 4, 5, 6)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 7)
	fmt.Println(s, len(s), cap(s))
}
