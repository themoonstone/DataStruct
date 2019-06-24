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
type LinkedList struct {
	Head 	*Node	// 链表头节点
	size	int		// 链表大小
}
// 链表的构造函数(参数包含节点和元素)
func (node *Node) ConstructorAll(next *Node, element interface{})  {
	node.Element = element
	node.Next = next
}
// 链表的构造函数(参数包含元素)
func (node *Node) ConstructorWithData( element interface{})  {
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

// 获取大小
func (list *LinkedList) GetSize() int {
	return list.size
}

// 判断链表是否为空
func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// 从链表头部插入数据
func (list *LinkedList) InsertHead(element interface{}) {
	// 通过新元素生成一个node,让该node的next指向head节点，然后使该node成为新的头节点
	var node *Node = new(Node)
	node.ConstructorAll(list.Head, element)
	list.Head = node
	list.size++
}
// 从链表指定位置插入数据
// 需要判断index是否为0，如果为0则说明实际上是从链表头部插入数据
func (list *LinkedList) Insert(index int, element interface{})  {
	// index合法性判断
	if index < 0 || index > list.size {
			panic("insert failed! illegal index")
	}
	if index == 0{
		list.InsertHead(element)
	} else {
		// 从指定位置插入，先给一个prev节点,从head开始
		prev := list.Head
		for i := 0; i < index - 1; i++ {
			// 找index的前一个节点，此处正好找到index-1处的节点
			prev = prev.Next
		}
		var node *Node = new(Node)
		node.Element = element
		node.Next = prev.Next 	// 新节点的next指向原prev节点的next
		prev.Next = node		// 新节点node作为原prev节点的next，注意，与上面的顺序不能错
		list.size++				// 链表大小加1
	}
}
// 从链表尾部插入数据
func (list *LinkedList) InsertTail(element interface{}) {
	//// 想办法先得到最后的索引
	//prev := list.Head
	//for i := 0; i < list.size - 1; i++ {
	//	prev = prev.Next
	//}
	//prev.Next = nil
	//prev.Element = element
	//list.size++
	// 我特么真是脑子有坑，已经实现了insert居然还XJB写了上面的逻辑
	list.Insert(list.size, element)
}