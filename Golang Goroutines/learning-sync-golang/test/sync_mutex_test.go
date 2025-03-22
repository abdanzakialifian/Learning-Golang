package test

import (
	"fmt"
	"sync"
	"sync/atomic"
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

	str1 := "Abdan"
	str2 := "Zaki"
	str3 := "Alifian"

	pool.Put(&str1)
	pool.Put(&str2)
	pool.Put(&str3)

	for range 10 {
		wait.Add(1)
		go func() {
			data := pool.Get().(*string)
			fmt.Println(*data)
			pool.Put(data)
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("Complete.")
}

func TestMap(t *testing.T) {
	data := sync.Map{}
	wait := sync.WaitGroup{}

	var addToMap = func(value int) {
		defer wait.Done()
		data.Store(value, value)
	}

	for i := range 100 {
		wait.Add(1)
		go addToMap(i)
	}

	wait.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

	fmt.Println("Complete.")
}

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait()

	fmt.Println("Done ", value)

	cond.L.Unlock()
}

func TestCondSignal(t *testing.T) {
	for i := range 10 {
		group.Add(1)
		go WaitCondition(i + 1)
	}

	go func() {
		for range 10 {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait()
}

func TestCondBroadcast(t *testing.T) {
	for i := range 10 {
		group.Add(1)
		go WaitCondition(i + 1)
	}

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}

func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var counter int64 = 0

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", counter)
}
