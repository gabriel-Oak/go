package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5)
	slice1 = append(array1, 4, 5)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, array1)
	fmt.Println(slice2)
}
