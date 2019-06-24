package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type linkedList struct {
	head *Node
}

type Node struct {
	data Item
	next *Node
}

type Item generic.Type

func main() {
	var lList linkedList
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node3 := Node{3, nil}
	lList.head = &node1
	node1.next = &node2
	node2.next = &node3

	lList.printlist()
	lList.insertAtBegin(0)
	lList.insertAtEnd(4)
	lList.insertAtEnd(5)
	lList.printlist()
	rLList := lList.reverseList()
	rLList.printlist()
	lList.insertAtIndex(6, 5)
	lList.printlist()
}

func (ll *linkedList) printlist() {
	head := ll.head
	for head != nil {
		fmt.Print(head.data)
		head = head.next
	}
	fmt.Println()
}

func (ll *linkedList) insertAtBegin(data Item) {
	node := Node{data, nil}
	node.next = ll.head
	ll.head = &node
}

func (ll *linkedList) insertAtEnd(data Item) {
	head := ll.head
	node := Node{data, nil}
	for head.next != nil {
		head = head.next
	}
	head.next = &node
}

func (ll *linkedList) insertAtIndex(data Item, index int) {
	count := 0
	head := ll.head
	for head != nil {
		if index == count {
			if index == 0 {
				ll.insertAtBegin(data)
				break
			} else {
				node := Node{data, nil}
				node.next = head.next
				head.next = &node
			}
		}
		count++
		head = head.next
	}
	if index >= count || index < 0 {
		fmt.Println("Index out of bounds")
	}
}

func (ll *linkedList) reverseList() (rLList linkedList) {
	head := ll.head
	for head != nil {
		rLList.insertAtBegin(head.data)
		head = head.next
	}
	return rLList
}
