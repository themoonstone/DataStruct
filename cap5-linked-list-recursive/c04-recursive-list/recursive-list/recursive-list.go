package recursive_list

// leetcode中的listNode定义
type ListNode struct {
	Val int
	Next *ListNode
}

// 删除所有指定元素(用递归的方式)
func removeElements(head *ListNode, val int) *ListNode {
	// 将一个链表分为头节点+由后面节点组成的链表两个部分
	if head == nil {
		return nil
	}
	// res表示删除指定值之后的链表
	res := removeElements(head.Next, val)
	if head.Val == val {
		// 这一步是进行删除的关键步骤
		return res
	} else {
		head.Next = res
	}
	return head
}