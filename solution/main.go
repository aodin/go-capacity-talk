package main

import "fmt"

func main() {
	var array [2]int
	slice := array[:1]
	fmt.Println(slice, len(slice), cap(slice))

	var outputs [][]int

	for i := 0; i < 3; i += 1 {
		output := make([]int, len(slice))
		copy(output, slice)
		outputs = append(outputs, append(output, i))
	}
	fmt.Println(outputs)
}
