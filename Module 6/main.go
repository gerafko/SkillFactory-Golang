package main

import (
	"fmt"
)

func main() {
	var str string
	var num1, num2 float64
	fmt.Scanln(&num1)
	fmt.Scanln(&str)
	fmt.Scanln(&num2)
	switch str {
	case "+":
		fmt.Printf("Результат суммирования: %F", num1+num2)
	case "-":
		fmt.Printf("Результат вычетания: %F", num1-num2)
	case "*":
		fmt.Printf("Результат умножения: %F", num1*num2)
	case ":":
		if num2 == 0 {
			fmt.Printf("Деление на ноль - ни ни")
		} else {
			fmt.Printf("Результат деления: %F", num1/num2)
		}
	}
}
