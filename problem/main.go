package main

import "fmt"

func main() {
	var array [1]int // Switch between [1]int and [2]int
	slice := array[:1]
	fmt.Println(slice, len(slice), cap(slice))

	var outputs [][]int

	for i := 0; i < 3; i += 1 {
		outputs = append(outputs, append(slice, i))
	}
	fmt.Println(outputs)
}
