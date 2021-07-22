package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 500
}

func deposit(amount int ,wg *sync.WaitGroup) {
	mutex.Lock()
	/*
		fmt.Println("Enter cash u want to deposit: ")

		var amount int
		fmt.Scanf("%d",&amount)

	*/
	fmt.Printf("Depositing %d to account with initial balance: %d\n", amount, balance)
	balance += amount
	fmt.Println("Money deposited and your present bal: ",balance)

	mutex.Unlock()
	wg.Done()
}

func withdraw(amount int ,wg *sync.WaitGroup) {
	mutex.Lock()
	/*
		fmt.Println("Enter cash u want to withdraw: ")
		var amount int
		fmt.Scanf("%d",&amount)

	*/
	fmt.Printf("Withdrawing %d from account with present balance: %d\n", amount, balance)
	if balance -amount <0 {
		fmt.Println("Insufficient balance .Your acc bal: " ,balance)
		mutex.Unlock()
		//time.Sleep(time.Millisecond)
		//wg.Done()
	}else {
		balance -= amount
		fmt.Println("Your curr bal: ",balance)
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	fmt.Println("Bank Account Processor")

	var wg sync.WaitGroup
	rand.Seed(time.Now().UTC().UnixNano())
	for i:=0;i<1;i++ {
		wg.Add(1)
		go deposit(rand.Intn(900)+100,&wg)
	}
	for j:=0;j<2;j++{
		wg.Add(1)
		go withdraw(rand.Intn(1500)+100,&wg)
	}
	wg.Wait()



	fmt.Printf("New Balance %d\n", balance)
}