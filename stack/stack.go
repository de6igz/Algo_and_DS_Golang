package main

import "fmt"

type Stack struct {
	items []interface{}
}

func (stack *Stack) Push(item interface{}) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() {
	index := len(stack.items) - 1
	fmt.Println(stack.items[index])
	stack.items = stack.items[:index]

}

func (stack *Stack) Peek() {
	if stack.isEmpty() == false {
		index := len(stack.items) - 1
		fmt.Println(stack.items[index])
	}

}

func (stack *Stack) isEmpty() bool {
	if len(stack.items) == 0 {
		fmt.Println("Stack Is Empty")
	}
	return true
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Peek()
}
