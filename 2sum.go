package main

import (
	"fmt"
	"sort"
)

func main() {
	// define sample array
	array := []int{7,3,4,1,8,11}
	arrM := make(map[int]int)
	for j,value := range array{
	 
	arrM[value]=j
//	fmt.Println(j,value)
	}
//	fmt.Println("arrM",arrM)
	// define target
	target:= 11
	// sort the array
	sort.Ints(array)
	// check if max number is greater than target
	for _,value := range array {
	  // fmt.Println(target,value)
	if(arrM[target-value] > 0){
	   fmt.Printf("index1: %d, index2: %d \n", arrM[value], arrM[target-value])
	}
	}
	// print varilable
//	fmt.Println(array, target)
	

}
