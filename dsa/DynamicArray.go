package main

import "fmt"

type SliceArray struct {
	Data     []int
	length   int
	capacity int
}

func (a *SliceArray) AddAtBegin(value int) {
	// fmt.Println("a.length == a.capacity ", a.length, a.capacity)
	if a.length == a.capacity {
		newCapacity := a.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		// fmt.Println("newCapacity ", newCapacity, a.length)
		// newArray := newArray(newCapacity)
		newArray := make([]int, newCapacity)
		for i := 0; i < a.length; i++ {
			// fmt.Println("a.Data[i] ", a.Data[i])
			newArray[i+1] = a.Data[i]
		}
		// fmt.Println("newArray, newArray.capacity", newArray, newCapacity)
		a.Data = newArray
		a.capacity = newCapacity
		// fmt.Println("a.Data, a.capacity, a.length ", a.Data, a.capacity, a.length)
	} else {
		count := a.length
		for x := count; x > 0; x-- {
			a.Data[x] = a.Data[x-1]
		}
	}
	a.Data[0] = value
	a.length++

}

func (a *SliceArray) DisplayAll() []int {
	for i := 0; i < a.length; i++ {
		fmt.Printf("value at Index %d is %d \n", i, a.Data[i])
	}

	return a.Data[0:a.length]

}

func (a *SliceArray) AddAtEnd(value int) {
	if a.length == a.capacity {
		newCapacity := a.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		newArray := make([]int, newCapacity)
		for i := 0; i < a.length; i++ {
			// fmt.Println("a.Data[i] ", a.Data[i])
			newArray[i] = a.Data[i]
		}
		a.Data = newArray
		a.capacity = newCapacity
		// fmt.Println(a.Data, a.capacity, a.length)
	}
	a.Data[a.length] = value
	a.length++
}

func (a *SliceArray) RemoveElementFrom(index int) ([]int, error) {
	count := a.length
	// if index != 0 {
	// 	count = index + 1
	// }
	for i := 0; i < count; i++ {
		if i >= index {
			a.Data[i] = a.Data[i+1]
		}
	}
	return a.Data, fmt.Errorf("err")
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
