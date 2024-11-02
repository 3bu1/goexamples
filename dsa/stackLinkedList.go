package main

import "fmt"

type StackLinkedListack struct {
	top *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (Stl *StackLinkedListack) push(value int) {
	NewNode := &Node{Value: value, Next: Stl.top}
	Stl.top = NewNode
}

func (Stl *StackLinkedListack) pop() int {
	value := Stl.top.Value
	Stl.top = Stl.top.Next
	return value
}

func (Stl *StackLinkedListack) DisplayLL() {
	for Stl.top.Next != nil {
		fmt.Println("Stl.top ", Stl.top.Value)
		Stl.top = Stl.top.Next
	}
	fmt.Println("Stl.top ", Stl.top.Value)
}
