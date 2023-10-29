package main

import "fmt"

func main() {
	array := []int{10, 20, 30, 40}

	math := make(chan int)

	go sum(array, math)

	fmt.Println(<-math)
}

func sum(intArray []int, ch chan int) {
	result := 0

	for _, val := range intArray {
		result += val
	}

	ch <- result
}
