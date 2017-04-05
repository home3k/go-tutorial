package main

import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int)  {
	mu.Lock()
	balance += amount
	defer mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	defer mu.Unlock()
	return b
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}
