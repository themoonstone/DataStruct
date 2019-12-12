package main

import (
	"fmt"
	"themoonstone/DataStruct/cap3-stack-queue/c02-brackets-match/arrayStack"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
/*
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。
*/
func isValid(s string) bool {
	var as arrayStack.ArrayStack
	as.Array.InitDefaultArray()
	for _, schar := range s {
		if string(schar) == "{" || string(schar) == "[" || string(schar) == "(" {
			as.Push(string(schar))
		}
		if string(schar) == "}" && as.Pop() != "{" {
			return false
		}
		if string(schar) == "]" && as.Pop() != "[" {
			return false
		}
		if string(schar) == ")" && as.Pop() != "(" {
			return false
		}
	}

	return as.IsEmpty()
}

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

	var brackets string = "{[}()[]{}"
	fmt.Println(isValid(brackets))
}
