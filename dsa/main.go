package main

func main() {
	// aa := &Array{capacity: 10}
	// aa.AddAtBegin(11)
	// aa.AddAtBegin(12)
	// aa.AddAtBegin(13)
	// aa.AddAtEnd(14)
	// aa.AddAtEnd(15)
	// aa.AddAtBegin(16)
	// // aa.AddAtBegin(16)
	// aa.DisplayAll()
	// sa := &SliceArray{capacity: 0}
	// sa.AddAtEnd(12)
	// sa.AddAtBegin(11)
	// sa.AddAtBegin(14)
	// sa.AddAtBegin(15)
	// sa.AddAtBegin(13)
	// sa.RemoveElementFrom(4)
	// sa.DisplayAll()
	// st := &Stack{element: make([]int, 0)}
	// st.push(1990)
	// st.push(1991)
	// st.push(1992)
	// fmt.Println("st.pop() ", st.pop())
	// st.displayAll()
	st := &StackLinkedListack{}
	st.push(10)
	st.push(20)
	st.push(30)
	st.push(40)
	st.pop()
	st.pop()
	st.DisplayLL()
}
