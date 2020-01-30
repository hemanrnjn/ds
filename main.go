package main

import (
	"fmt"

	ll "github.com/hemanrnjn/ds/linkedlist"
	"github.com/hemanrnjn/ds/stack"
)

func main() {
	linkedListOps()
	stackOps()
}

func linkedListOps() {
	lList := ll.NewLList()
	node1 := ll.NewNode(1)
	node2 := ll.NewNode(2)
	node3 := ll.NewNode(3)
	lList.Head = &node1
	node1.Next = &node2
	node2.Next = &node3

	lList.Printlist()
	lList.InsertAtBegin(0)
	lList.InsertAtEnd(4)
	lList.InsertAtEnd(5)
	lList.Printlist()
	rLList := lList.ReverseList()
	rLList.Printlist()
	lList.InsertAtIndex(6, 5)
	lList.Printlist()
	fmt.Println(lList.PrintFromNthNode(3))
}

func stackOps() {
	stack := stack.NewStack()
	stack.PrintStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.PrintStack()
	stack.Pop()
	stack.Pop()
	stack.PrintStack()
}
