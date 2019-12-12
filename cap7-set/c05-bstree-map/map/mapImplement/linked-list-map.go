package mapImplement

import (
	"fmt"
	"log"
	"themoonstone/DataStruct/cap7-set/c05-bstree-map/utils/interfaces"
)

// 基于链表实现map

// 基本结构
// map节点基本结构
type LNode struct {
	Key   interface{}
	Value interface{}
	Next  *LNode
}

// map结构
type LMap struct {
	// 虚拟头节点
	DummyHead *LNode
	size      int // 节点数量
}

// node节点初始化(无参)
func (ln *LNode) Constructor() {
	ln.Key = nil
	ln.Value = nil
	ln.Next = nil
}

// 节点的构造函数(参数包含节点和元素)
func (ln *LNode) ConstructorAll(next *LNode, key, value interface{}) {
	ln.Key = key
	ln.Value = value
	ln.Next = next
}

// node节点初始化(有参)
func (ln *LNode) ConstructorWithArgs(key, value interface{}) {
	ln.Key = key
	ln.Value = value
	ln.Next = nil
}

// 链表初始化
func (lm *LMap) Constructor() {
	lm.DummyHead = new(LNode)
	lm.DummyHead.Constructor()
	lm.size = 0
}

// 获取指定key值的node信息(辅助函数)
func (lm *LMap) getNode(key interface{}) *LNode {
	// 遍历获取
	var node = lm.DummyHead
	for node.Next != nil {
		if interfaces.Compare(node.Next.Key, key) == 0 {
			return node.Next
		}
		node = node.Next
	}
	return nil
}

func (lm *LMap) Add(key interface{}, value interface{}) {
	node := lm.getNode(key)
	if node != nil {
		// 引用修改，也可以修改原值
		node.Value = value
	} else {
		// 在链表头插入元素
		var lnode *LNode = new(LNode)
		// 新的头节点的next指向原来的实际头节点
		lnode.ConstructorAll(lm.DummyHead.Next, key, value)
		// node成为新的实际头节点
		lm.DummyHead.Next = lnode
		lm.size++
	}
}

// 删除
func (lm *LMap) Remove(key interface{}) interface{} {
	if lm.getNode(key) == nil {
		log.Panicf("key [%v] is not exists\n", key)
	}
	var dummyNode = lm.DummyHead
	for dummyNode.Next != nil {
		if interfaces.Compare(dummyNode.Next.Key, key) == 0 {
			// 删除
			// 要删除的节点
			var delNode *LNode = dummyNode.Next
			dummyNode.Next = delNode.Next
			delNode.Next = nil
			lm.size--
			return delNode.Value
		}
		dummyNode = dummyNode.Next
	}
	return nil
}

func (lm *LMap) Set(key, value interface{}) {
	if lm.getNode(key) == nil {
		log.Panicf("the key [%v] is not exist\n", key)
	}
	// 遍历获取
	var node = lm.DummyHead
	for node.Next != nil {
		if interfaces.Compare(node.Next.Key, key) == 0 {
			node.Next.Value = value
			return
		}
		node = node.Next
	}
}

func (lm *LMap) Get(key interface{}) interface{} {

	node := lm.getNode(key)
	if node == nil {
		return nil
	} else {
		return node.Value
	}
}

func (lm *LMap) Contains(key interface{}) bool {
	return lm.getNode(key) != nil
}
func (lm *LMap) IsEmpty() bool {
	return lm.size == 0
}
func (lm *LMap) Size() int {
	return lm.size
}

// 添加一个打印函数
// 以逗号分隔
func (lm *LMap) String() string {
	var res []string
	// 遍历
	var headNode = lm.DummyHead.Next
	if headNode != nil {
		res = append(res, fmt.Sprintf("<%v,%v>", headNode.Key, headNode.Value))
	}
	for headNode.Next != nil {
		res = append(res, fmt.Sprintf(",<%v,%v>", headNode.Next.Key, headNode.Next.Value))
		headNode = headNode.Next
	}
	return fmt.Sprintf("%v\n", res)
}
