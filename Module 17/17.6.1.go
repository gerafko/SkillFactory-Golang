package main

import (
	"fmt"
	"time"
)

func main() {
	// Текущее время
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			c2 <- 1
		}
	}()
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			c1 <- 1
		}
	}()
	var f, s int
	for i := 0; i < 100; i++ {
		select {
		case <-c1:
			fmt.Println("1.Получено сообщение из первого канала")
			f++
			fmt.Println(i)
		case <-c2:
			fmt.Printf("2.Получено сообщение из второго канала")
			s++
			fmt.Println(i)
		}
	}
	fmt.Println("Первый - ", f)
	fmt.Println("Первый - ", s)
}
