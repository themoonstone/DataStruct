package main

import "fmt"
import "themoonstone/DataStruct/cap2-array/c01-insert/array"
func main() {
	var a array.Array
	a.InitDefaultArray()

	a.AddLast(1)
	a.AddFirst(20)
	a.Insert(1,1)
	fmt.Println(a.GetSize())
	fmt.Printf("data : %d\n", a.data)
}

