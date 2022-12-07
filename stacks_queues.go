package main

import "fmt"

// Stack

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (result int) {
	lastIndex := len(s.items) - 1
	result = s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return
}

// Queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (result int) {
	result = q.items[0]
	q.items = q.items[1:]
	return
}

func main() {
	// Testing Stack
	fmt.Println("Stack:")
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	removedValue := myStack.Pop()
	fmt.Println("Removed", removedValue)
	fmt.Println(myStack)

	// Testing Queue
	fmt.Println("Queue:")
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	removedValue = myQueue.Dequeue()
	fmt.Println("Removed", removedValue)
	fmt.Println(myQueue)
}
