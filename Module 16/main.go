package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

type Client struct {
	amount int
	mutex  sync.RWMutex
}

func (c *Client) Deposit(amount int) {
	c.mutex.Lock()
	c.amount += amount
	//fmt.Println("Deposit: ", c.amount)
	c.mutex.Unlock()
}

func (c *Client) Withdrawal(amount int) (err string) {
	c.mutex.Lock()
	if c.amount < amount {
		return "Операция не может быть выполнена " + strconv.Itoa(c.amount) + "<" + strconv.Itoa(amount)
	}
	c.amount -= amount
	//fmt.Println("Withdrawal " + strconv.Itoa(c.amount))
	c.mutex.Unlock()
	return ""
}

func (c *Client) Balance() {
	c.mutex.RLock()
	fmt.Println("Balance: ", c.amount)
	c.mutex.RUnlock()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a Client
	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				for {
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-500)+500))
					a.Deposit(rand.Intn(9) + 1)
				}
			}()
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			go func() {
				for {
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-500)+500))
					a.Deposit(rand.Intn(4) + 1)
				}
			}()
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		switch s[0] {
		case "balance":
			a.Balance()
		case "deposit":
			amount, _ := strconv.Atoi(s[1])
			a.Deposit(amount)
		case "withdrawal":
			amount, _ := strconv.Atoi(s[1])
			er := a.Withdrawal(amount)
			if er != "" {
				fmt.Println(er)
			}
		case "exit":
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}

}
