package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	data Item
	Next *Node
}

type Item interface{}

func NewLList() LinkedList {
	var lList LinkedList
	return lList
}

func NewNode(data Item) Node {
	var node = Node{data, nil}
	return node
}

func (ll *LinkedList) Printlist() {
	head := ll.Head
	for head != nil {
		fmt.Print(head.data)
		head = head.Next
	}
	fmt.Println()
}

func (ll *LinkedList) InsertAtBegin(data Item) {
	node := Node{data, nil}
	node.Next = ll.Head
	ll.Head = &node
}

func (ll *LinkedList) InsertAtEnd(data Item) {
	head := ll.Head
	node := Node{data, nil}
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &node
}

func (ll *LinkedList) InsertAtIndex(data Item, index int) {
	count := 0
	Head := ll.Head
	for Head != nil {
		if index == count {
			if index == 0 {
				ll.InsertAtBegin(data)
				break
			} else {
				node := Node{data, nil}
				node.Next = Head.Next
				Head.Next = &node
			}
		}
		count++
		Head = Head.Next
	}
	if index >= count || index < 0 {
		fmt.Println("Index out of bounds")
	}
}

func (ll *LinkedList) ReverseList() *LinkedList {
	head := ll.Head
	curr := head
	var next *Node
	var prev *Node
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.Head = prev
	return ll
}

func (ll *LinkedList) PrintFromNthNode(n int) Item {
	head := ll.Head
	ref := ll.Head
	count := 0
	for count < n {
		ref = ref.Next
		count++
	}

	for ref.Next != nil {
		head = head.Next
		ref = ref.Next
	}

	return head.data
}
