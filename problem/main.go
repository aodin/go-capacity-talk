package main

import "fmt"

func main() {
	var array [2]int // Switch between [1]int and [2]int
	slice := array[:1]
	fmt.Println("Initial", slice, len(slice), cap(slice))

	var outputs [3][]int
	for i, _ := range outputs {
		outputs[i] = append(slice, i)
	}
	fmt.Println("Result:", outputs)
}
