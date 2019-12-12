package main

import "fmt"
import "themoonstone/DataStruct/cap2-array/c03-c-e-d/array"

func main() {
	var a array.Array
	a.InitDefaultArray()

	a.AddLast(1)
	a.AddFirst(20)
	a.Insert(1, 1)
	fmt.Println("size:", a.GetSize())
	fmt.Printf("data : %v\n", a.String())
	// 获取指定位置的元素
	fmt.Printf("before update val : %v\n", a.Get(0))
	// 设置指定位置的元素
	a.Set(2, 100)
	fmt.Printf("after update val : %v\n", a.String())
	a.Remove(0)
	fmt.Println("after delete val : ", a.String())
	a.RemoveHead()
	fmt.Println("after delete head : ", a.String())
	a.RemoveTail()
	fmt.Println("after delete tail : ", a.String())
}
