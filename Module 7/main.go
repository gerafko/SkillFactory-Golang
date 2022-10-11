package main

import (
	"calc"
	"fmt"
)

func main() {
	var str string
	var num1, num2 float64
	fmt.Scanln(&num1)
	fmt.Scanln(&str)
	fmt.Scanln(&num2)
	arifm := calc.NewCalculator(num1, num2, str)
	fmt.Print(arifm.Calculate())
}
