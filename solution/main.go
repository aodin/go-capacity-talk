package main

import "fmt"

func main() {
	var array [2]int
	slice := array[:1]
	fmt.Println("Initial", slice, len(slice), cap(slice))

	var outputs [3][]int
	for i, _ := range outputs {
		outputs[i] = make([]int, len(slice))
		copy(outputs[i], slice)
		outputs[i] = append(outputs[i], i)
	}
	fmt.Println("Result:", outputs)
}
