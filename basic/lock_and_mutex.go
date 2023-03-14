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

func (a *Account) getBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) withDrawBalance(wdAmount int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if wdAmount > a.balance {
		pl("Not enough amount")
	} else {
		a.balance -= wdAmount

		pl("Balance withdrawn")
	}
}

func main() {

	var acc Account
	acc.balance = 100

	pl("Balance : ", acc.getBalance())

	for i := 1; i <= 12; i++ {
		go acc.withDrawBalance(10)
	}

	time.Sleep(2 * time.Millisecond)

}
