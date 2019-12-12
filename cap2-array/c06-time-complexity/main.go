package main

import (
	"fmt"
	"themoonstone/DataStruct/cap2-array/c05-dynamic-array/array"

	"themoonstone/DataStruct/cap2-array/c05-dynamic-array/Student"
)

func main() {
	Int()
	Stu()
}

func Int() {
	var a array.Array
	a.InitDefaultArray()
	a.AddFirst(100)
	fmt.Printf("before resize data : %v\n", a.String())
	for i := 0; i < 20; i++ {
		a.Insert(i, i)
	}
	fmt.Println("size:", a.GetSize())
	fmt.Printf("data : %v\n", a.String())
	// 获取指定位置的元素
	fmt.Printf("before update val : %v\n", a.Get(0))
	// 设置指定位置的元素
	a.Set(1, 100)
	fmt.Printf("after update val : %v\n", a.String())
	a.Remove(0)
	fmt.Println("after delete val : ", a.String())
	//a.RemoveHead()
	//fmt.Println("after delete head : ", a.String())
	a.RemoveTail()
	fmt.Println("after delete tail : ", a.String())
}

func Stu() {
	var as array.Array
	as.InitDefaultArray()
	as.AddLast(Student.Student{"troytan", 120})
	as.AddLast(Student.Student{"xiaomei", 90})
	as.AddLast(Student.Student{"alice", 105})
	fmt.Println(as.String())
}
