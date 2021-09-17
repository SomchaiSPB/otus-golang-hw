package main

import (
	"fmt"
)

type IntStack struct {
	arr []int
}

func (i *IntStack) Pop() int {
	last := i.arr[len(i.arr) - 1]
	i.arr = i.arr[:len(i.arr) - 1]

	return last
}

func (i *IntStack) Push(num int) {
	i.arr = append(i.arr, num)
}

func main() {
	var s IntStack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())
}
