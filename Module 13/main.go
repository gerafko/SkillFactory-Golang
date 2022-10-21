package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (t *BinaryTreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.val == value {
		fmt.Println("This node value already exists ", value)
		return errors.New("This node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &BinaryTreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &BinaryTreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *BinaryTreeNode) PrintInorder() {
	if t == nil {
		return
	}
	fmt.Print(t.val, "    ")
	t.left.PrintInorder()
	t.right.PrintInorder()
}

func (t *BinaryTreeNode) LookUp(value int) {
	if t.val == value {
		fmt.Println("Founded:", t.val, "    ", &t)
		return
	}
	if t.val != value && t.left == nil && t.right == nil {
		fmt.Println("LookUp - ", value, " not exist")
		return
	}
	if t.left != nil {
		t.left.LookUp(value)
	}
	if t.right != nil {
		t.right.LookUp(value)
	}
}

func WideLookUp(graph [][]int, start, val int) {
	data := make([][]int, len(graph))[len(graph):]
	data = append(data, graph[start])

	used := make(map[int]bool, len(graph))
	used = map[int]bool{start: true}

	for len(data) > 0 {
		for _, key := range data[len(data)-1] {
			if _, ok := used[key]; !ok {
				data = append(data, graph[key])
			}
			used[key] = true
			if _, ok := used[val]; ok {
				fmt.Println("Неориентированный граф - поиск в ширину, закончен. Значение", val, " найдено")
				return
			}
		}
		data = data[:len(data)-1]
	}
	fmt.Println("!Неориентированный граф - поиск в ширину, закончен. Значение", val, " НЕ найдено")
}

func OrientalFastWay(data [][][]int, val int) {
	fastWay := 1000000
	for _, key := range data {
		sum := 0
		for _, key1 := range key {
			sum += key1[1]
			if key1[0] == val && sum < fastWay {

				fastWay = sum
			}
		}
	}
	if fastWay == 0 {
		fmt.Println("НЕ найдено")
		return
	}
	fmt.Println("Быстрый путь: ", fastWay)
}

func main() {
	fmt.Println("Бинарное дерево")
	binary := BinaryTreeNode{val: 5}
	look := 0
	for i := 0; i < 10; i++ {
		look = rand.Intn(10)
		binary.Insert(look)
	}
	binary.PrintInorder()
	fmt.Println()
	binary.LookUp(100)
	binary.LookUp(look)
	fmt.Println()
	fmt.Println()
	fmt.Println("Не ориентированный граф")
	graph := [][]int{
		{1, 2},
		{0, 2},
		{0, 1, 3, 4, 5},
		{2},
		{2, 6},
		{2, 6},
		{4, 5},
	}
	WideLookUp(graph, 2, 6)
	WideLookUp(graph, 2, 100)
	fmt.Println()
	fmt.Println()
	fmt.Println("Ориентированный граф")
	orientGraph := [][][]int{
		{{1, 1}, {3, 6}},
		{{0, 1}, {2, 4}, {3, 3}, {4, 9}},
		{{0, 6}, {1, 4}},
		{{1, 3}, {4, 2}},
	}
	OrientalFastWay(orientGraph, 4)
}
