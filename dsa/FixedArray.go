package main

import "fmt"

type Array struct {
	Data     []int
	length   int
	capacity int
}

func newArray(capacity int) *Array {
	return &Array{Data: make([]int, 0), length: 0, capacity: capacity}
}

func (a *Array) AddAtBegin(value int) {
	if a.Data == nil {
		a.Data = make([]int, a.capacity)
	}
	count := a.length
	for x := count; x > 0; x-- {
		a.Data[x] = a.Data[x-1]
	}
	a.Data[0] = value
	a.length++
}

func (a *Array) DisplayAll() []int {
	for i := 0; i < a.length; i++ {
		fmt.Printf("value at Index %d is %d \n", i, a.Data[i])
	}

	return a.Data[0:a.length]

}

func (a *Array) AddAtEnd(value int) {
	if a.Data == nil {
		a.Data = make([]int, a.capacity)
	}
	a.Data[a.length] = value
	a.length++
}

// func main() {
// aa := &Array{capacity: 10}
// aa.AddAtBegin(11)
// aa.AddAtBegin(12)
// aa.AddAtBegin(13)
// aa.AddAtEnd(14)
// aa.AddAtEnd(15)
// // aa.AddAtBegin(16)
// aa.DisplayAll()
// }
