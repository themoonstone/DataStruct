package recursive_list

import "fmt"

// leetcode中的listNode定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除所有指定元素(用递归的方式)
func RemoveElements(head *ListNode, val int, depth int) *ListNode {
	// 打印递归深度
	fmt.Print(generateDepth(depth))
	fmt.Printf("Call : remove %d in %v\n", val, head)
	// 将一个链表分为头节点+由后面节点组成的链表两个部分
	if head == nil {
		fmt.Print(generateDepth(depth))
		fmt.Printf("Return : %v\n", head)
		return nil
	}

	// res表示删除指定值之后的链表
	res := RemoveElements(head.Next, val, depth+1)
	fmt.Print(generateDepth(depth))
	fmt.Printf("After : remove %d in %v\n", val, res)

	var ret *ListNode = new(ListNode)
	if head.Val == val {
		// 这一步是进行删除的关键步骤
		ret = res
	} else {
		head.Next = res
		ret = head
	}
	fmt.Print(generateDepth(depth))
	fmt.Printf("Return: %v\n", ret)
	return ret
}

// 生成代表递归深度的符号
func generateDepth(depth int) string {
	var res string
	for i := 0; i < depth; i++ {
		res += "--"
	}
	return res
}

// 格式化输出
func (ln *ListNode) String() string {
	var res []interface{}
	// 输出当前链表
	for cur := ln; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
		res = append(res, "->")
	}
	res = append(res, "NULL")
	return fmt.Sprintf("Head %v Tail\n", res)
}

// 生成链表
func (ln *ListNode) New(nums []int) {
	ln.Val = nums[0]
	cur := ln
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil} // ln.Next = cur.Next
		cur = cur.Next                     // 进行这一步之后，cur后移，但ln以及ln的Next不变
	}
}
