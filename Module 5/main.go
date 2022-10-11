package main

import "fmt"

func main() {
	const (
		_ = iota
		jan
		feb
		mar
		apr
		may
	)
	fmt.Println(jan)
	fmt.Println(feb)
	fmt.Println(mar)
	fmt.Println(apr)
	fmt.Println(may)
}
