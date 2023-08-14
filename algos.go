package main

//
//import (
//	"fmt"
//)
//
//// Node Определение структуры для элемента связного списка
//type Node struct {
//	data interface{}
//	next *Node
//}
//
//// LinkedList Определение структуры для связного списка
//type LinkedList struct {
//	head *Node
//}
//
//// Append Метод для добавления элемента в конец списка
//func (ll *LinkedList) Append(data interface{}) {
//	newNode := &Node{data: data, next: nil}
//	if ll.head == nil {
//		ll.head = newNode
//		return
//	}
//
//	current := ll.head
//	for current.next != nil {
//		current = current.next
//	}
//	current.next = newNode
//}
//
//// Remove Метод для удаления элемента из списка по значению
//func (ll *LinkedList) Remove(data interface{}) {
//	if ll.head == nil {
//		return
//	}
//
//	if ll.head.data == data {
//		ll.head = ll.head.next
//		return
//	}
//
//	current := ll.head
//	for current.next != nil {
//		if current.next.data == data {
//			current.next = current.next.next
//			return
//		}
//		current = current.next
//	}
//}
//
//// Display Метод для вывода элементов списка
//func (ll *LinkedList) Display() {
//	current := ll.head
//	for current != nil {
//		fmt.Println(current.data)
//		current = current.next
//	}
//}
//
//func main() {
//	// Создание нового связного списка
//	linkedList := LinkedList{}
//
//	// Добавление элементов в список
//	linkedList.Append("123")
//	linkedList.Append(2)
//	linkedList.Append(3)
//	linkedList.Append(4)
//	for i := 10; i < 99; i++ {
//		linkedList.Append(i)
//	}
//
//	// Вывод элементов списка
//	fmt.Println("Linked List:")
//	linkedList.Display()
//
//	// Удаление элемента
//	linkedList.Remove(2)
//
//	// Вывод элементов после удаления
//	fmt.Println("Linked List after removal:")
//	linkedList.Display()
//}
