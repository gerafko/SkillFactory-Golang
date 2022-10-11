package calc

import "fmt"

type calculator struct {
	oper   string  // operator
	n1, n2 float64 // numbers
}

func NewCalculator(n1 float64, n2 float64, oper string) calculator {
	return calculator{n1: n1, n2: n2, oper: oper}
}

func (a calculator) Calculate() (res float64) {
	switch a.oper {
	case "+":
		res = sum(a.n1, a.n2)
	case "-":
		res = sub(a.n1, a.n2)
	case "*":
		res = mult(a.n1, a.n2)
	case ":":
		if a.n2 == 0 {
			fmt.Printf("Деление на ноль - ни ни")
		} else {
			res = div(a.n1, a.n2)
		}
	default:
		fmt.Println("Ты что-то не то ввел")
	}
	return
}

func sum(n1, n2 float64) float64 {
	return n1 + n2
}

func mult(n1, n2 float64) float64 {
	return n1 * n2
}

func sub(n1, n2 float64) float64 {
	return n1 - n2
}

func div(n1, n2 float64) float64 {
	return n1 / n2
}
