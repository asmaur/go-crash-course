package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) WithDraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough")
	} else {
		fmt.Printf("%d withdrawn. Balance: %d\n", v, a.balance)
		a.balance -= v
	}
}

func main() {
	var account Account
	account.balance = 1200
	pl("Balance: ", account.GetBalance())

	for i := 0; i < 12; i++ {
		go account.WithDraw(10)
	}
	time.Sleep(2 * time.Second)
}
