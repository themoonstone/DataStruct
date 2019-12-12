package main

import (
	"fmt"
	"themoonstone/DataStruct/cap4-linked-list/c06-list-stack/arrayStack"
	"themoonstone/DataStruct/cap4-linked-list/c06-list-stack/listStack"
	"themoonstone/DataStruct/cap4-linked-list/c06-list-stack/stack"
	"time"
)

func main() {
	//var ls listStack.ListStack
	//ls.Constructor()
	//// 入栈
	//for i := 0; i < 5; i++ {
	//	ls.Push(i)
	//	fmt.Println(ls.String())
	//}
	//
	//// 出栈
	//ls.Pop()
	//fmt.Println(ls.String())

	// ListStack
	var ls listStack.ListStack
	ls.Constructor()
	fmt.Println(opStack(&ls, 10000000))
	// ArrayStack
	var as arrayStack.ArrayStack
	as.Constructor()
	fmt.Println(opStack(&as, 10000000))
	// 通过比较，可以发现其实使用数组栈的时间还要短一些，但这个结果并不确定。
	// 在有些机器上，或者使用其它语言实现，也会出现链表栈耗时短的情况。
	// 主要是因为链表栈在插入的时候会不断的去查找并new一个新的内存空间，而通常这一步比较消耗时间
	// 但通过这个比较，我们可以看到链表栈与数组栈的时间复杂度是同一级别的，不会有非常大的差距
}

// 性能比较
// 分别为ArrayStack和ListStack添加100000条数据，然后再全部出栈，比较运行时间
func opStack(stack stack.Stack, opCount int) int64 {
	start := time.Now().UnixNano()
	for i := 0; i < opCount; i++ {
		stack.Push(i)
	}
	for i := 0; i < opCount; i++ {
		stack.Pop()
	}
	return time.Now().UnixNano() - start
}
