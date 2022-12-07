package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
	// My implementation:
	// n.next = l.head
	// l.head = n
	// l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	fmt.Print(l)
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")

	// My implementation:
	// for i, node := 0, l.head; i < l.length; i++ {
	// 	fmt.Println(node.data)
	// 	node = node.next
	// }
}

func (l *linkedList) deleteWithValue(value int) {
	// My implementation
	// previous := l.head
	// toCheck := l.head
	// if toCheck.data == value {
	// 	l.head = toCheck.next
	// 	l.length--
	// } else {
	// 	for toCheck.data != value && toCheck.next != nil {
	// 		previous = toCheck
	// 		toCheck = toCheck.next
	// 	}
	// 	if toCheck.data == value {
	// 		previous.next = toCheck.next
	// 		l.length--
	// 	}
	// }

	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}

		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 26}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	fmt.Println(myList)
	myList.printListData()
	myList.deleteWithValue(48)
	myList.deleteWithValue(47)
	myList.printListData()

	emptyList := linkedList{}
	emptyList.deleteWithValue(2)
}
