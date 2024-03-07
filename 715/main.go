package main

import "fmt"

/*
Determine whether a doubly linked list is a palindrome. What if it’s singly linked?

For example, 1 -> 4 -> 3 -> 4 -> 1 returns True while 1 -> 4 returns False.
*/

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Add(data ...int) {
	for _, item := range data {
		node := Node{data: item}

		// check if it's the first element
		if dll.head == nil {
			dll.head = &node
			dll.tail = &node
		} else {
			// append to the list
			node.prev = dll.tail
			dll.tail.next = &node
			dll.tail = &node
		}
	}
}

func (dll *DoublyLinkedList) Print() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) IsPalindrom() bool {
	start := dll.head
	end := dll.tail
	for {
		if start == end {
			// last element from iteration
			return true
		}
		if start.data == end.data {
			// step - by - step
			start = start.next
			end = end.prev
		} else {
			// data differs
			return false
		}
	}
}

func main() {
	// 1 -> 4 -> 3 -> 4 -> 1
	dll := NewDoublyLinkedList()
	dll.Add(1, 4, 3, 4, 1)
	dll.Print()

	if dll.IsPalindrom() {
		fmt.Println("List is palindrom")
	} else {
		fmt.Println("List is not a palindrom")
	}

	dll = NewDoublyLinkedList()
	dll.Add(1, 4)
	dll.Print()

	if dll.IsPalindrom() {
		fmt.Println("List is palindrom")
	} else {
		fmt.Println("List is not a palindrom")
	}
}
