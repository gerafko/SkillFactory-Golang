package main

import (
	"fmt"
	"time"
)

func main() {
	// Текущее время
	c1 := make(chan int)
	c2 := make(chan int)
	start := time.Now()

	var f, s int

	for i := 0; i < 100; i++ {
		select {
		case <-c1:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("1.Получено сообщение из первого канала")
			f++
			fmt.Println(i)
		case <-c2:
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("2.Получено сообщение из второго канала")
			s++
			fmt.Println(i)
		default:
			fmt.Printf("ВРЕМЯ %v\n", time.Since(start))
		}
	}
	//fmt.Println("Первый - ", f)
	//fmt.Println("Второй - ", s)
}
