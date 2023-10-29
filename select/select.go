package main

import (
	"fmt"
	"time"
)

func main() {
	green := make(chan string)
	yellow := make(chan string)
	red := make(chan string)

	go greenCh(green)
	go yellowCh(yellow)
	go redCh(red)

	for {
		select {
		case msg := <-green:
			fmt.Println(msg)
		case msg := <-yellow:
			fmt.Println(msg)
		case msg := <-red:
			fmt.Println(msg)
		}
	}
}

func greenCh(green chan string) {
	for {
		time.Sleep(1 * time.Second)
		green <- "green"
	}
}

func yellowCh(yellow chan string) {
	for {
		time.Sleep(3 * time.Second)
		yellow <- "yellow"
	}
}

func redCh(red chan string) {
	for {
		time.Sleep(1 * time.Second)
		red <- "red"
	}
}
