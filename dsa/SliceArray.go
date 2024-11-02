// package dsa

// import "fmt"

// // MySlice is a custom slice struct
// type MySlice struct {
// 	array    []int // underlying array
// 	length   int   // current length
// 	capacity int   // total capacity
// }

// // NewSlice initializes a new slice with a given capacity.
// func NewSlice(capacity int) *MySlice {
// 	return &MySlice{
// 		array:    make([]int, capacity),
// 		length:   0,
// 		capacity: capacity,
// 	}
// }

// // Append adds an element to the slice, resizing if necessary.
// func (s *MySlice) Append(value int) {
// 	if s.length == s.capacity {
// 		newCapacity := s.capacity * 2
// 		if newCapacity == 0 {
// 			newCapacity = 1
// 		}
// 		newArray := make([]int, newCapacity)
// 		n := len(s.array)
// 		if len(newArray) < n {
// 			n = len(newArray)
// 		}

// 		for i := 0; i < n; i++ {
// 			newArray[i] = s.array[i]
// 		}
// 		s.array = newArray
// 		s.capacity = newCapacity
// 	}
// 	s.array[s.length] = value
// 	s.length++
// }

// // Get retrieves the element at a specific index
// func (s *MySlice) Get(index int) (int, error) {
// 	if index < 0 || index >= s.length {
// 		return 0, fmt.Errorf("index out of bounds")
// 	}
// 	return s.array[index], nil
// }

// // Set assigns a value to a specific index
// func (s *MySlice) Set(index int, value int) error {
// 	if index < 0 || index >= s.length {
// 		return fmt.Errorf("index out of bounds")
// 	}
// 	s.array[index] = value
// 	return nil
// }

// // Len returns the current length of the slice
// func (s *MySlice) Len() int {
// 	return s.length
// }

// // Cap returns the capacity of the slice
// func (s *MySlice) Cap() int {
// 	return s.capacity
// }
// func OperateSlice() {
// 	fmt.Println("in operatesice")

// 	s := NewSlice(2)
// 	fmt.Println("Initial capacity:", s.Cap())

// 	s.Append(10)
// 	s.Append(20)
// 	fmt.Println("Length after adding two elements:", s.Len())
// 	fmt.Println("Capacity after adding two elements:", s.Cap())

// 	s.Append(30)
// 	fmt.Println("Length after adding third element:", s.Len())
// 	fmt.Println("Capacity after adding third element:", s.Cap())
// }
