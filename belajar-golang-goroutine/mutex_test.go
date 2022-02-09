package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("total balance", account.GetBalance())
}

//deadlock
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

//end deadlock
//testing deadlock
func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "eko",
		Balance: 10,
	}
	user2 := UserBalance{
		Name:    "eko",
		Balance: 10,
	}

	go Transfer(&user1, &user2, 1)
	go Transfer(&user2, &user1, 2)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balane ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balane ", user2.Balance)
}

//end testing deadlock