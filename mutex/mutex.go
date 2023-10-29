package main

import (
	"fmt"
	"sync"
	"time"
)

var accountBalance int = 0

var wg sync.WaitGroup

var mutex sync.Mutex

func transaction(amount int, action string, name string) {
	if action == "CREDIT" {
		accountBalance += amount

		fmt.Printf("$%d Amount has been credited to your account: %s\n", amount, name)
	}

	if action == "DEBIT" {
		if (accountBalance - amount) < 0 {
			fmt.Printf("Insufficient funds (%d) to debit $%d from your account: %s\n", accountBalance, amount, name)
			return
		}

		mutex.Lock()

		accountBalance -= amount

		fmt.Printf("$%d Amount has been debited from your account: %s\n", amount, name)

		mutex.Unlock()
	}
}

func accountHolder1(name string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		transaction(10, "CREDIT", name)
		time.Sleep(time.Millisecond * 500)
	}
}

func accountHolder2(name string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		transaction(10, "DEBIT", name)
		time.Sleep(time.Millisecond * 500)
	}
}

func accountHolder3(name string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		transaction(1, "CREDIT", name)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	wg.Add(3)

	go accountHolder1("Jazz")
	go accountHolder2("Jane")
	go accountHolder3("John")

	wg.Wait()

	fmt.Printf("Final account balance is: %d\n", accountBalance)
}
