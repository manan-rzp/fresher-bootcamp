package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type Account struct {
	balance int
}

func (c *Account) Withdraw(amount int, mu *sync.Mutex) {
	mu.Lock()

	if c.balance >= amount {
		c.balance -= amount
	}else{
		panic("Low Balance")
	}

	mu.Unlock()
}
func (c *Account) Deposit(amount int, mu *sync.Mutex) {
	mu.Lock()

	c.balance += amount

	mu.Unlock()
}
func main() {
	var mu sync.Mutex

	c := Account{balance: 100}

	for i := 0; i < 500; i++ {

		if rand.Intn(10) > 5 {
			go c.Withdraw(rand.Intn(2000), &mu)
		}else{
			go c.Deposit(rand.Intn(500), &mu)
		}
		fmt.Println(c.balance)
	}

	time.Sleep(time.Second)

	fmt.Println(c.balance)

}