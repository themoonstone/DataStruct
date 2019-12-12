package main

import (
	"fmt"
	"themoonstone/DataStruct/cap5-linked-list-recursive/c02-remove-val-local/list"
)

func main() {
	// 给定一个数组，生成一个链表
	var nums []int = []int{1, 2, 6, 3, 4, 5, 6}
	var listNode *list.ListNode = new(list.ListNode)

	listNode.New(nums)
	fmt.Println(listNode.String())

	fmt.Println(list.RemoveElementsWithDummyHead(listNode, 6).String())
}
