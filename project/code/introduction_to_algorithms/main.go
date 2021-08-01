package main

import (
	"code_test/project/code/introduction_to_algorithms/sort"
	"fmt"
)

func main(){
	//test insertion sort

	arrInsertionSort := []int{5,2,4,6,1,3}

	fmt.Println("insertion soort",sort.InsertionSort(arrInsertionSort))
}
