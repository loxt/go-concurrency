package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChannel := make(chan string, 3)

	bufferedChannel <- "first"
	bufferedChannel <- "second"
	bufferedChannel <- "third"

	time.Sleep(3 * time.Second)

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}
