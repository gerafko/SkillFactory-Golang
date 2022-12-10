package main

import (
	"fmt"
	"sync"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 10

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var counter int64 = 0
	increment := func() {
		counter += step
		c.Signal()
	}
	// Не всегда вычисление этой переменной будет приводить к верному
	// результату в счётчике, но для правильных значений
	// и для удобства - можно
	var iterationCount int = int(endCounterValue / step)
	c.L.Lock()
	for i := 1; i <= iterationCount; i++ {
		go increment()
		c.Wait()
	}
	// Ожидаем поступления сигнала
	c.L.Unlock()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
