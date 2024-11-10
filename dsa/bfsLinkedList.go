package main

import (
	"fmt"
)

// TreeNode represents a node in the tree
type TreeNode struct {
	Value    int
	Children []*TreeNode
}

// BFS performs a breadth-first search on the tree and prints each node's value
func BFS(root *TreeNode) {
	if root == nil {
		return
	}

	// Initialize a queue and enqueue the root node
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Dequeue a node from the front of the queue
		current := queue[0]
		queue = queue[1:]

		// Process the current node (for example, print its value)
		fmt.Println(current.Value)

		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
}

func DFS(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	// loop stack and check children
	for len(stack) > 0 {
		current := stack[0]
		fmt.Println("current left :", current.Value)
		stack = stack[1:]
		for _, child := range current.Children {
			stack
		}
	}
}
