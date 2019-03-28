package main

import (
	"fmt"
	"themoonstone/DataStruct/cap5-linked-list-recursive/c03-recursive/recursive-array"
)

func main() {
	var s recursive_array.RecursiveArray
	var array []int = []int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("sum : %d\n", s.Sum(array))
}
