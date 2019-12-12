package main

import "fmt"
import "themoonstone/DataStruct/cap2-array/c02-query-update/array"

func main() {
	var a array.Array
	a.InitDefaultArray()

	a.AddLast(1)
	a.AddFirst(20)
	a.Insert(1, 1)
	fmt.Println(a.GetSize())
	fmt.Printf("data : %d\n", a)
	// 获取指定位置的元素
	fmt.Printf("before update val : %d\n", a.Get(0))
	// 设置指定位置的元素
	a.Set(2, 100)
	fmt.Printf("after update val : %d\n", a.Get(0))
}
