package list

// leetcode中的listNode定义
 type ListNode struct {
 	Val int
 	Next *ListNode
 }

// 删除所有指定元素
func removeElements(head *ListNode, val int) *ListNode {
	// 判断头节点是否满足删除条件(有可能会连续满足条件)
	// 注意，head是否为空必须每次都进行判断，因为可能每次移动之后的head都为空
	for head != nil && head.Val == val {
		var delNode = head
		head = head.Next
		delNode.Next = nil	// 和链表断掉联系
	}
	// 有可能全部节点都包含了指定元素，那样的话就会连续删除所有节点，所以需要再次判断
	if head == nil {
		return nil
	}

	// 从中间位置开始删除
	// 找到待删除节点的前一个节点
	prev := head	// 前面的代码运行完成之后，已经可以确保head不包含指定元素
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil	// 和链表断掉联系
		} else {
			// 元素不匹配，prev后移
			prev = prev.Next
		}
	}
	return head
}