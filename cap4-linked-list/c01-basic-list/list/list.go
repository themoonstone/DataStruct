package list

import "fmt"

// 链表的节点结构
type Node struct {
	// 当前节点所存储的数据
	Element interface{}
	// 指向下一个节点的指针(引用)
	Next *Node
}

// 链表的构造函数(参数包含节点和元素)
func (node *Node) ConstructorAll(next *Node, element interface{}) {
	node.Element = element
	node.Next = next
}

// 链表的构造函数(参数包含元素)
func (node *Node) ConstructorWithData(element interface{}) {
	node.Element = element
	node.Next = nil
}

// 链表的构造函数(无参)
func (node *Node) Constructor() {
	node.Element = nil
	node.Next = nil
}

// 实现链表的格式化输出
func (node *Node) String() string {
	return fmt.Sprintln(node.Element)
}
