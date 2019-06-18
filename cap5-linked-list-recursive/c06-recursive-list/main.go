package main

import (
	"fmt"
	"themoonstone/DataStruct/cap5-linked-list-recursive/c05-recursive-list-print/recursive-list"
)

func main() {
	// 通过数组生成新的链表
	var num []int = []int{1,2,6,3,4,5,6}
	var ln *recursive_list.ListNode = new(recursive_list.ListNode)
	// 生成新的链表
	ln.New(num)
	fmt.Println(ln.String())

	fmt.Println(recursive_list.RemoveElements(ln, 6, 0))
}
