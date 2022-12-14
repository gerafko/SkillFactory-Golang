package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	bankAccount int
}

const amountOfTransactions int = 1000

func doTransaction(summ int, bank chan<- int, transactionDone <-chan int, wg *sync.WaitGroup) int {
	defer wg.Done()
	bank <- summ
	return <-transactionDone
}

func (b Bank) createBank(wg *sync.WaitGroup) (bank chan int, transactionDone chan int) {
	bank = make(chan int)
	transactionDone = make(chan int)
	go func() {
		defer wg.Done()
		for summ := range bank {
			b.bankAccount += summ
			transactionDone <- b.bankAccount
		}
	}()
	return
}

func main() {
	var b Bank
	var amountOfGoRoutines int = amountOfTransactions * 2
	var wg sync.WaitGroup
	wg.Add(amountOfGoRoutines)
	bank, transactionDone := b.createBank(&wg)
	for i := 1; i <= amountOfTransactions; i++ {
		go doTransaction(1, bank, transactionDone, &wg)
		go doTransaction(-1, bank, transactionDone, &wg)
	}
	wg.Wait()
	wg.Add(1)
	close(bank)
	wg.Wait()
	fmt.Println(b.bankAccount)
}
