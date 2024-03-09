package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

func printTo15(){
	for i := 1; i <= 15; i++ {
		pl("Func 1 :", i)
	}
}

func printTo10(){
	for i := 1; i <= 10; i++ {
		pl("Func 2 :", i)
	}
}

func nums1(channel chan int){
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int){
	channel <- 4
	channel <- 5
	channel <- 6
}

type Account struct {
	balance int
	lock sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int){
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance{
		pl("Not Enough Money in Account")
	} else {
		a.balance -= v
		fmt.Printf("%d Withdrawn: Balance :%d\n", v, a.balance)
	}
}

func main(){
	//go printTo15()
	//go printTo10()
	//time.Sleep(20 & time.Second)

	//channel1 := make(chan int)
	//channel2 := make(chan int)
	//go nums1(channel1)
	//go nums2(channel2)
	//pl(<-channel1)
	//pl(<-channel2)
	//pl(<-channel1)
	//pl(<-channel2)

	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}
