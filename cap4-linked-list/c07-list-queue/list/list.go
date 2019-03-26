package list

import "fmt"

// 链表的节点结构
type Node struct {
	// 当前节点所存储的数据
	Element 	interface{}
	// 指向下一个节点的指针(引用)
	Next 		*Node
}

// 链表基本结构
type LinkedListQueue struct {
	Head 	*Node	// 链表头节点
	Tail	*Node	// 链表尾节点
	size	int		// 链表大小
}
// 节点的构造函数(参数包含节点和元素)
func (node *Node) ConstructorAll(next *Node, element interface{})  {
	node.Element = element
	node.Next = next
}
// 节点的构造函数(参数包含元素)
func (node *Node) ConstructorWithData( element interface{})  {
	node.Element = element
	node.Next = nil
}
// 节点的构造函数(无参)
func (node *Node) Constructor() {
	node.Element = nil
	node.Next = nil
}
// 链表的构造函数(针对虚拟头节点有用)
func (lq *LinkedListQueue) Constructor() {

	lq.Head = nil
	lq.Tail = nil
	lq.size = 0
}

// 获取大小
func (lq *LinkedListQueue) GetSize() int {
	return lq.size
}

// 判断链表是否为空
func (lq *LinkedListQueue) IsEmpty() bool {
	return lq.size == 0
}

// 入队
func (lq *LinkedListQueue) Enqueue(element interface{})  {
	var node Node
	node.ConstructorWithData(element)
	// 生成一个节点，让队尾next指向该节点
	// 首先注意tail为空的情况 ，tail为空说明当前的链表队列为空

	if lq.Tail == nil {
		lq.Tail = &node
		lq.Head = lq.Tail
	}else {
		lq.Tail.Next = &node
		// 尾部节点记得要向后移
		lq.Tail = lq.Tail.Next
	}
	lq.size++
}

// 出队
// 注意两个问题
// 1. 先判断队列是否为空
// 2. 在执行完了删除操作之后，需要再判断一下队列是否为空
func (lq *LinkedListQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		panic("Queue is empty!")
	}
	// head节点为队头，出队
	var delNode *Node = new(Node)
	delNode = lq.Head
	// 头节点后移
	lq.Head = lq.Head.Next
	delNode.Next = nil
	if lq.Head == nil {
		// 在队列中只有一个元素的情况下，删除该节点之后队列为空，需要重新维护
		lq.Tail = nil
	}
	lq.size--
	return delNode.Element
}

// 获取队头元素
// 先判断是否为空
func (lq *LinkedListQueue) GetFront() interface{} {
	if lq.IsEmpty() {
		panic("Queue is empty!")
	}
	return lq.Head.Element
}

// 格式化输出
func (list *LinkedListQueue) String() string {
	var res []interface{}
	res = append(res, "Queue:Head [")
	for cur := list.Head; cur != nil; cur = cur.Next {
		res = append(res, cur.Element)
		res = append(res, "->")
	}
	res = append(res, "NULL] Tail")

	return fmt.Sprintf("%v\n", res)
}