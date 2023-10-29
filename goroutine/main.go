package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printSomething("One", &wg)
	go printSomething("Two", &wg)

	wg.Wait()
}

func printSomething(value string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 500)
	}

	defer wg.Done()
}
