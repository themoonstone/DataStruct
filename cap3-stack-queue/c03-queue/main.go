package main

import (
	"fmt"
	"themoonstone/DataStruct/cap3-stack-queue/c03-queue/ArrayQueue"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
/*
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。
*/

func main() {
	var aq ArrayQueue.ArrayQueue
	aq.Constructor()
	// 入栈操作
	for i := 0; i < 20; i++ {
		aq.Enqueue(i)
		fmt.Println(aq.String())
		if i%3 == 0 {
			aq.Dequeue()
		}
	}
	fmt.Println(aq.String())
}
