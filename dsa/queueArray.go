// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type queue struct {
	elements []int
	front    int
	rear     int
}

func (q *queue) enqueue(value int) {
	fmt.Println(q.rear)
	if q.rear == -1 {
		q.front = 0
		q.rear = 0
	}

	q.elements[q.rear] = value

	fmt.Println(q.elements, q.rear, value)
	q.rear++

	if q.front == q.rear {
		q.front = 0
		q.rear = 0
	}
}

func (q *queue) dequeue() {

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	}
	fmt.Println(q.elements[q.front])
	if q.front != -1 {
		q.front++
	}
}

func (q *queue) display() {
	for a := q.front; a < q.rear; a++ {
		fmt.Println("a is ", q.elements[a])
	}
}

// func main() {
// 	fmt.Println("Hello, 世界")

// 	qu := &queue{elements: make([]int, 10)}
// 	qu.enqueue(6)
// 	qu.enqueue(7)
// 	qu.enqueue(8)
// 	qu.display()
// 	qu.dequeue()
// 	qu.display()
// 	fmt.Println()

// }
