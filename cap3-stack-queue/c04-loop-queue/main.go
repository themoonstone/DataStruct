package main

import (
	"fmt"
	"themoonstone/DataStruct/cap3-stack-queue/c04-loop-queue/LoopQueue"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
/*
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。
*/

func main() {
	var queue LoopQueue.LoopQueue
	queue.Constructor()
	// 入栈操作
	for i := 0 ; i < 20; i++ {
		queue.Enqueue(i)
		fmt.Printf("i : %d -- queue : %v\n",i, queue.String())
		if i % 3 == 0 {
			queue.Dequeue()
		}
	}
	fmt.Println("----------------------")
	fmt.Println(queue.String())
}
