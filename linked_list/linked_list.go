package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	data interface{}
	next *Node
}

func (l *LinkedList) Append(data interface{}) {
	nextNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = nextNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = nextNode

}
func (l *LinkedList) Display() {
	current := l.head
	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println(current.data)
	return
}

func (l *LinkedList) Remove(data interface{}) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next

	}
}
func main() {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append("hello")
	list.Display()
	list.Remove(3)
	list.Remove(1)
	fmt.Println("List after removal ")
	list.Display()

}
