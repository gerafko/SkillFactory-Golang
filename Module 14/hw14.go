package main

import (
	"fmt"
)

func main() {
	fmt.Print("First array size:")
	var arrSize1, arrSize2 int
	fmt.Scanln(&arrSize1)
	fmt.Print("Second array size: ")
	fmt.Scanln(&arrSize2)
	fmt.Println("Enter first array: ")
	arr1 := make([]int, arrSize1)
	arr2 := make([]int, arrSize2)
	val := 0
	for i := 0; i < arrSize1; i++ {
		fmt.Scanln(&val)
		arr1[i] = val
	}
	fmt.Println("Enter second array:")
	for i := 0; i < arrSize2; i++ {
		fmt.Scanln(&val)
		arr2[i] = val
	}
	crossMap := make(map[int]int, arrSize1)
	for i := 0; i < arrSize1; i++ {
		for j := 0; j < arrSize2; j++ {
			if arr1[i] == arr2[j] {
				crossMap[i] = arr1[i]
			}

		}
	}
	fmt.Println(crossMap)
}
