package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
}

func TestSyncRWMutex(t *testing.T) {
	account := BankAccount{}
	for range 100 {
		go func() {
			for range 100 {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance : ", account.GetBalance())
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(balance int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + balance
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Abdan",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Zaki",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User", user1.Name, "with Balance", user1.Balance)
	fmt.Println("User", user2.Name, "with Balance", user2.Balance)
}

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

func (user *UserBalance) Change(balance int) {
	user.Balance += balance
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1 : ", user1.Name)
	user1.Change(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 : ", user2.Name)
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}

	for range 100 {
		go RunAsynchronous(&group)
	}

	group.Wait()
	fmt.Println("Complete.")
}

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for range 100 {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", counter)
}

var counter = 0

func OnlyOnce() {
	counter++
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}

	wait := sync.WaitGroup{}

	pool.Put("Abdan")
	pool.Put("Zaki")
	pool.Put("Alifian")

	for range 10 {
		wait.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("Complete.")
}
