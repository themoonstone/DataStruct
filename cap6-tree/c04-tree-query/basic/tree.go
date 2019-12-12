package basic

import (
	"themoonstone/DataStruct/cap6-tree/utils/interfaces"
)

type BasicTree struct {
	Root *Node
	size int
}

type Node struct {
	Left    *Node
	Right   *Node
	Element interface{}
}

// 初始化
func (tree *BasicTree) Constructor(e interface{}) {
	tree.size = 0
	tree.Root = nil
}

// 生成一个新的节点
func generate(element interface{}) *Node {
	return &Node{
		Left:    nil,
		Right:   nil,
		Element: element,
	}
}

// 大小
func (tree *BasicTree) Size() int {
	return tree.size
}

// 是否为空
func (tree *BasicTree) IsEmpty() bool {
	return tree.size == 0
}

// 插入方法
func (tree *BasicTree) Add(e interface{}) {

	// 调用递归函数
	tree.Root = tree.add(tree.Root, e)
}

// 以当前node节点为根节点，进行插入
// 返回插入新节点后二分搜索树的根
func (tree *BasicTree) add(node *Node, e interface{}) *Node {
	// 递归终止条件
	// 如果node为空，一定会new一个新的节点
	if node == nil {
		tree.size++
		return generate(e)
	}
	// 如果node不为空，进行判断
	if interfaces.Compare(e, node.Element) == -1 {
		node.Left = tree.add(node.Left, e)
	} else if interfaces.Compare(e, node.Element) == 1 {
		node.Right = tree.add(node.Right, e)
	}
	return node

}

// 查询二分搜索树中是否包含某个元素(采用递归的方式)
func (tree *BasicTree) Contains(e interface{}) bool {
	return tree.contains(tree.Root, e)
}

// 查询递归内部函数
func (tree *BasicTree) contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}
	if interfaces.Compare(e, node.Element) == 0 {
		return true
	}
	// 在左边
	if interfaces.Compare(e, node.Element) == -1 {
		return tree.contains(node.Left, e)
	} else {
		return tree.contains(node.Right, e)
	}
}
