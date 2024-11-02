package main

import "fmt"

type Stack struct {
	element []int
}

func (st *Stack) push(value int) {
	st.element = append(st.element, value)
}

func (st *Stack) pop() int {
	value := 0
	if len(st.element) != 0 {
		value = st.element[len(st.element)-1]
		// st.element = st.element[0 : len(st.element)-1]
		st.element[len(st.element)-1] = 0
	}
	return value
}

func (st *Stack) displayAll() {
	fmt.Println("value: ", st.element)
}
