package data

import "fmt"

type Stack struct {
	Top *node
}

func (stack *Stack) Push(value interface{}) {
	if stack.Top == nil {
		stack.Top = &node{Data: value, Next: nil}
		return
	}

	stack.Top = &node{Data: value, Next: stack.Top}
}

func (stack *Stack) Pop() {
	stack.Top = stack.Top.Next
}

func (stack *Stack) Display() {

	s := stack.Top

	for s != nil {

		fmt.Println(s.Data)

		s = s.Next
	}

}
