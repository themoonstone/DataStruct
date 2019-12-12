package list

import "fmt"

// 链表的节点结构
type Node struct {
	// 当前节点所存储的数据
	Element interface{}
	// 指向下一个节点的指针(引用)
	Next *Node
}

// 链表基本结构
type LinkedList struct {
	DummyHead *Node // 链表头节点(修改为虚拟头节点)
	size      int   // 链表大小
}

// 节点的构造函数(参数包含节点和元素)
func (node *Node) ConstructorAll(next *Node, element interface{}) {
	node.Element = element
	node.Next = next
}

// 节点的构造函数(参数包含元素)
func (node *Node) ConstructorWithData(element interface{}) {
	node.Element = element
	node.Next = nil
}

// 节点的构造函数(无参)
func (node *Node) Constructor() {
	node.Element = nil
	node.Next = nil
}

// 链表的构造函数(针对虚拟头节点有用)
func (list *LinkedList) Constructor() {
	// 生成一个虚拟头节点,其元素为空，next也为空
	var dummyHead *Node = new(Node)
	dummyHead.Constructor()
	list.DummyHead = dummyHead
	list.size = 0
}

// 获取大小
func (list *LinkedList) GetSize() int {
	return list.size
}

// 判断链表是否为空
func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// 从链表指定位置插入数据
// 插入节点最重要的是找到待插入索引位置的前一个位置的节点
func (list *LinkedList) Insert(index int, element interface{}) {
	// index合法性判断
	if index < 0 || index > list.size {
		panic("insert failed! illegal index")
	}
	// 从指定位置插入，先给一个prev节点,从DummyHead开始
	prev := list.DummyHead
	// 注意：添加了dummyHead节点之后，初始的prev，也就是dummyHead指向的是索引为0的节点的前一个节点
	// 而在没有添加dummyHead之前，初始的prev指向的是索引为0的节点
	// 也就是说，现在我们的目的是要找到index索引位置的元素之前的节点，所以，只需要遍历到index就好了
	for i := 0; i < index; i++ {
		// 找index的前一个节点，此处正好找到index-1处的节点
		prev = prev.Next
	}
	var node *Node = new(Node)
	node.Element = element
	node.Next = prev.Next // 新节点的next指向原prev节点的next
	prev.Next = node      // 新节点node作为原prev节点的next，注意，与上面的顺序不能错
	list.size++           // 链表大小加1
}

// 从链表头部插入数据
func (list *LinkedList) InsertHead(element interface{}) {
	// 通过新元素生成一个node,让该node的next指向head节点，然后使该node成为新的头节点
	//var node *Node = new(Node)
	////node.ConstructorWithData(element)
	////node.Next = list.Head
	//node.ConstructorAll(list.Head, element)
	//list.Head = node

	list.Insert(0, element)
	//list.size++
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

// 实现查询功能
// 查询指定位置的元素
func (list *LinkedList) Get(index int) interface{} {
	// 首先判断索引的合法性
	if index < 0 || index > list.GetSize() {
		panic("illegal index")
	}
	// 从虚拟节点开始
	cur := list.DummyHead.Next
	// 遍历
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Element
}

// 查询头节点的元素
func (list *LinkedList) GetFirst() interface{} {
	return list.Get(0)
}

// 查询尾节点的元素
func (list *LinkedList) GetLast() interface{} {
	return list.Get(list.GetSize())
}

// 查询链表中是否包含了某个节点的元素
func (list *LinkedList) Contains(element interface{}) bool {
	//cur := list.DummyHead.Next
	//for i := 0; i < list.GetSize();i++ {
	//	if cur.Element == element {
	//		return true
	//	}
	//	cur = cur.Next
	//}
	// 更好的方法
	for cur := list.DummyHead.Next; cur != nil; cur = cur.Next {
		if cur.Element == element {
			return true
		}
	}
	return false
}

// 修改
func (list *LinkedList) Update(index int, element interface{}) {
	// 首先判断索引的合法性
	if index < 0 || index > list.GetSize() {
		panic("illegal index")
	}
	// 从虚拟节点开始
	cur := list.DummyHead.Next
	// 遍历
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Element = element
}

// 实现链表的格式化输出
// 需要知道想输出一个什么样的格式
// node->node->null
func (list *LinkedList) String() string {
	var res []interface{}
	for cur := list.DummyHead.Next; cur != nil; cur = cur.Next {
		res = append(res, cur.Element)
		res = append(res, "->")
	}
	res = append(res, "NULL")
	return fmt.Sprintf("%v\n", res)
}

// 删除链表
func (list *LinkedList) Delete(index int) interface{} {
	if index < 0 || index > list.GetSize() {
		panic("illegal index")
	}
	delNode := list.DummyHead
	// 找到要删除的节点的前一个节点
	for i := 0; i < index; i++ {
		delNode = delNode.Next // 这里满足条件的其实是要删除的节点的前一个节点
	}
	// 要删除的元素
	element := delNode.Next.Element
	delNode.Next = delNode.Next.Next
	delNode = nil
	list.size--
	return element
}

// 删除表头
func (list *LinkedList) DeleteHead() interface{} {
	return list.Delete(0)
}

// 删除表尾
func (list *LinkedList) DeleteTail() interface{} {
	return list.Delete(list.GetSize() - 1)
}

// 删除指定元素的节点(默认删除第一个)
//func (list *LinkedList) RemoveElement(e interface{}) {
//	// 遍历链表，遇到与e相等的节点即删除
//	// 定义一个虚拟头节点
//	VHead := list.DummyHead
//	prev := VHead
//	curNode := VHead.Next
//
//	// 遍历
//	for curNode != nil {
//		// 判断
//		if curNode.Element == e {
//			prev.Next = curNode.Next
//			curNode = curNode.Next
//			list.size--
//		} else {
//			prev = prev.Next
//			curNode = curNode.Next
//		}
//	}
//}

// 删除指定元素的节点
func (list *LinkedList) RemoveElement(e interface{}) {
	// 遍历链表，遇到与e相等的节点即删除
	// 定义一个虚拟头节点
	prev := list.DummyHead

	// 遍历
	for prev.Next != nil {
		// 判断
		if prev.Next.Element == e {
			// 要删除的节点
			toDelNode := prev.Next
			// 这里已经将toDelNode从链表中删除了
			prev.Next = toDelNode.Next
			// 这里用toDelNode.Next=nil是为了让toDelNode不再指向链表
			toDelNode.Next = nil
			list.size--
		} else {
			prev = prev.Next
		}
	}
}
