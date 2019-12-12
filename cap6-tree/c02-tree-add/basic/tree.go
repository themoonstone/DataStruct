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
	// 当前二叉树为空
	if tree == nil {
		tree.Root = generate(e)
		tree.size++
		return
	}
	// 调用递归函数
	tree.add(tree.Root, e)
}

// 以当前node节点为根节点，进行插入
func (tree *BasicTree) add(node *Node, e interface{}) {
	// 递归终止条件
	if interfaces.Compare(e, node.Element) == 0 {
		return
	} else if interfaces.Compare(e, node.Element) == -1 && node.Left == nil {
		tree.size++
		// 待插入元素比当前节点元素小，插到左边
		node.Left = generate(e)
		tree.size++
		return
	} else if interfaces.Compare(e, node.Element) == 1 && node.Right == nil {
		tree.size++
		// 待插入元素比当前节点元素大，插到右边
		node.Right = generate(e)
		tree.size++
		return
	} else if interfaces.Compare(e, node.Element) == -1 {
		// 缩减规模
		tree.add(node.Left, e)
	} else {
		tree.add(node.Right, e)
	}
}
