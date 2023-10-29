package main

import (
	"fmt"
	"time"
)

func main() {
	externalService1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		externalService1 <- "response from external service 1"
	}()

	select {
	case res := <-externalService1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	externalService2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		externalService2 <- "response from external service 1"
	}()

	select {
	case res := <-externalService2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout 2")
	}
}
