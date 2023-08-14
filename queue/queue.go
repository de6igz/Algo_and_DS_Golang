package main

import "fmt"

type Queue struct {
	elements []interface{}
}

func (queue *Queue) Insert(element interface{}) {
	queue.elements = append(queue.elements, element)
}
func (queue *Queue) Peek() interface{} {
	if len(queue.elements) == 0 {
		fmt.Println("Очердь пуста")
		return nil
	}
	return queue.elements[0]
}

func (queue *Queue) Remove() interface{} {
	if len(queue.elements) == 0 {
		fmt.Println("Очердь пуста")
		return nil
	}
	value := queue.elements[0]
	queue.elements = queue.elements[1:]
	return value
}

func (queue *Queue) Display() {
	fmt.Println(queue.elements)
}

func main() {
	queue := Queue{}
	queue.Insert(1)
	queue.Insert(2)

	fmt.Println(queue.Peek())
	queue.Remove()
	queue.Display()
	fmt.Println(queue.Peek())

}
