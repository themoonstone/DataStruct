package list

// 使用虚拟头节点

// 删除所有指定元素
func removeElementsWithDummyHead(head *ListNode, val int) *ListNode {
	// 给定一个虚拟头节点作为Head的前一个节点
	dummyHead := new(ListNode)
	dummyHead.Next = head
	prev := dummyHead
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
	return dummyHead.Next
}