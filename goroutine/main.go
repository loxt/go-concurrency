package main

import (
	"fmt"
	"time"
)

func main() {
	go printSomething("One")
	printSomething("Two")
}

func printSomething(value string) {
	for {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 500)
	}
}
