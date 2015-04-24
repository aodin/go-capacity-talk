package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 2, 3)
	fmt.Println(slice, len(slice), cap(slice))
}
