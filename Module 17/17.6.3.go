package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	go func() {
		for i := 1; i < 101; i++ {
			c1 <- i
		}
	}()
	go func() {
		for {

			<-c1
			fmt.Println(<-c1)
		}
	}()
	time.Sleep(1000 * time.Millisecond)
}
