package main

import (
	"fmt"
	"themoonstone/DataStruct/cap4-linked-list/c07-list-queue/list"
)

func main() {
	var lq list.LinkedListQueue
	lq.Constructor()
	// 入栈操作
	for i := 0 ; i < 5; i++ {
		lq.Enqueue(i)
		fmt.Println(lq.String())
		if i % 3 == 0 {
			lq.Dequeue()
		}
	}
	fmt.Println(lq.GetSize())
	fmt.Println(lq.String())
}

