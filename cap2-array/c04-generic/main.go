package main

import (
	"fmt"
	"themoonstone/DataStruct/cap2-array/c04-generic/Student"
	"themoonstone/DataStruct/cap2-array/c04-generic/array"
)
func main() {
	Int()
	Stu()
}

func Int()  {
	var a array.Array
	a.InitDefaultArray()

	a.AddLast("ls")
	a.AddFirst(20)
	a.Insert(1,1)
	fmt.Println("size:", a.GetSize())
	fmt.Printf("data : %v\n", a.String())
	// 获取指定位置的元素
	fmt.Printf("before update val : %v\n", a.Get(0))
	// 设置指定位置的元素
	a.Set(2,100)
	fmt.Printf("after update val : %v\n", a.String())
	a.Remove(0)
	fmt.Println("after delete val : ", a.String())
	a.RemoveHead()
	fmt.Println("after delete head : ", a.String())
	a.RemoveTail()
	fmt.Println("after delete tail : ", a.String())
}

func Stu() {
	var as array.Array
	as.InitDefaultArray()
	as.AddLast(Student.Student{"troytan",120})
	as.AddLast(Student.Student{"xiaomei",90})
	as.AddLast(Student.Student{"alice",105})
	fmt.Println(as.String())
}