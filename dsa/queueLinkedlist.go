package main

import (
	"fmt"
)

type LNode struct {
	value int
	Next  *LNode
}

type Lqueue struct {
	front *LNode
	rear  *LNode
}

func (q *Lqueue) enLqueue(value int) {
	newLNode := &LNode{value: value}
	if q.rear == nil {
		q.rear = newLNode
		q.front = newLNode
	} else {
		q.rear.Next = newLNode
		q.rear = newLNode
	}
}

func (q *Lqueue) deLqueue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("Lqueue is empty")
	}
	value := q.front.value
	q.front = q.front.Next

	if q.front == nil {
		q.rear = nil
	}
	return value, nil
}

func (q *Lqueue) IsEmpty() bool {
	return q.front == nil
}

func (q *Lqueue) peek() (int, error) {
	//if q.front is nil
	if q.front == nil {
		return 0, fmt.Errorf("Lqueue is empty")
	}
	return q.front.value, nil
}
