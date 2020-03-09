package stack

import (
	"fmt"
)

type Stack struct {
	Items []Item
}

type Item interface{}

func NewStack() Stack {
	var stack Stack
	return stack
}

func (stack *Stack) Push(data Item) {
	stack.Items = append(stack.Items, data)
}

func (stack *Stack) Pop() Item {
	if len(stack.Items) > 0 {
		ele := stack.Items[len(stack.Items)-1]
		stack.Items = stack.Items[:len(stack.Items)-1]
		return ele
	}
	return nil
}

func (stack *Stack) PrintStack() {
	for i := len(stack.Items) - 1; i >= 0; i-- {
		fmt.Print(stack.Items[i])
	}
}
