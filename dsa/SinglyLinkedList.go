// package main

// import (
// 	"fmt"
// )

// // Node represents a node in the linked list
// type Node struct {
// 	data int
// 	next *Node
// }

// // LinkedList represents the linked list
// type LinkedList struct {
// 	head *Node
// }

// // InsertAtEnd adds a new node with given data at the end of the list
// func (ll *LinkedList) InsertAtEnd(data int) {
// 	newNode := &Node{data: data}
// 	if ll.head == nil {
// 		ll.head = newNode
// 	} else {
// 		current := ll.head
// 		for current.next != nil {
// 			current = current.next
// 		}
// 		current.next = newNode
// 	}
// }

// // InsertAtBeginning adds a new node with given data at the beginning of the list
// func (ll *LinkedList) InsertAtBeginning(data int) {
// 	newNode := &Node{data: data, next: ll.head}
// 	ll.head = newNode
// }

// // Delete deletes the first node with the given data
// func (ll *LinkedList) Delete(data int) {
// 	if ll.head == nil {
// 		return
// 	}
// 	if ll.head.data == data {
// 		ll.head = ll.head.next
// 		return
// 	}
// 	current := ll.head
// 	for current.next != nil && current.next.data != data {
// 		current = current.next
// 	}
// 	if current.next != nil {
// 		current.next = current.next.next
// 	}
// }

// // Display prints out the elements in the list
// func (ll *LinkedList) Display() {
// 	current := ll.head
// 	for current != nil {
// 		fmt.Printf("%d -> ", current.data)
// 		current = current.next
// 	}
// 	fmt.Println("nil")
// }

// func main() {
// 	ll := &LinkedList{}

// 	ll.InsertAtEnd(10)
// 	ll.InsertAtEnd(20)
// 	ll.InsertAtEnd(30)
// 	fmt.Print("After inserting at end: ")
// 	ll.Display()

// 	ll.InsertAtBeginning(5)
// 	fmt.Print("After inserting at beginning: ")
// 	ll.Display()

// 	ll.Delete(20)
// 	fmt.Print("After deleting 20: ")
// 	ll.Display()
// }
