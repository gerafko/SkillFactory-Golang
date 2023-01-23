package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	const N = 5
	const n = 10
	wg.Add(N)
	for i := 1; i <= N; i++ {
		go func(routine int) {
			for i := 0; i < n; i++ {
				fmt.Println(routine)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
