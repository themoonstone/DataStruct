package main

import (
	"fmt"
	"themoonstone/DataStruct/cap3-stack-queue/c01-stack-basic/arrayStack"
)

func main() {
	var as arrayStack.ArrayStack
	as.Array.InitDefaultArray()
	// 入栈操作
	for i := 0; i < 5; i++ {
		as.Push(i)
	}
	fmt.Println(as.String())
	// 出栈操作
	as.Pop()
	fmt.Println(as.String())
	// 查看栈顶
	fmt.Println(as.Peek())
	// 获取栈的大小
	fmt.Println(as.GetSize())
	// 获取栈的容积
	fmt.Println(as.GetCap())
	// 判断栈是否为空
	fmt.Println(as.IsEmpty())
}
