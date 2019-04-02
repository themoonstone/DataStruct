package bst

import (
	"themoonstone/DataStruct/cap6-binary-tree/utils/interfaces"
)

// 生成一个二分搜索树

// 节点
type Node struct {
	// 存储的元素
	Element 	interface{}
	// 左右子节点
	Left, Rigth	*Node
}
// 树
type BST struct {
	// 根节点
	Root	*Node
	// 存储元素数量
	Size	int
}

// 构造函数

// 节点
func (node *Node)Constructor(element interface{}) {
	node.Element = element
	node.Left = nil
	node.Rigth = nil
}

// 生成一个新的节点
func generate(element interface{}) *Node {
	return &Node{
		Left:nil,
		Rigth:nil,
		Element:element,
	}
}

// 树
func (bst *BST) Constructor() {
	bst.Root = nil
	bst.Size = 0
}

// 获取二分搜索树大小
func (bst *BST)GetSize() int {
	return bst.Size
}
// 判断二分搜索树是否为空
func (bst *BST) IsEmpty() bool {
	return bst.Size == 0
}

// 插入元素(非递归实现)
func (bst *BST) Insert(element interface{}) {
	if bst.Root == nil {
		// 生成根节点
		bst.Root.Constructor(element)
		bst.Size++
	}else {
		node := bst.Root
		for node != nil {
			// 传入的元素大于当前节点的元素，与当前节点的左子节点相比
			if interfaces.Compare(element, node.Element) == 1{
				node = node.Left
			} else {
				node = node.Rigth
			}
		}
		node.Constructor(element)
		// 大小加1
		bst.Size++
	}
}
// 插入元素，递归算法
func (bst *BST) Add(element interface{})  {
	if bst.Root == nil {
		bst.Root = generate(element)
		bst.Size++
		return
	}
	bst.add(bst.Root, element)
}

// 实际运行的递归算法
func (bst *BST) add(node *Node, element interface{}) {
	// 遵循递归两个条件
	// 1. 递归终止条件,节点的左子节点或右子节点为空
	if interfaces.Compare(element, node.Element) == 0 {
		return
	}
	// 待插入元素大于当前节点元素，且左子节点为空，可以直接加到左子节点处
	if interfaces.Compare(element , node.Element) == 1 && node.Left == nil {
		node.Left = generate(element)
		bst.Size++
	}else if interfaces.Compare(element, node.Element) == -1 && node.Rigth == nil {
		// 待插入元素小于当前节点元素，且右子节点为空，可以直接加到右子节点处
		node.Rigth = generate(element)
		bst.Size++
	}
	// 2. 缩减规模
	if interfaces.Compare(element , node.Element) == 1{
		// 与左子节点的元素进行比较
		bst.add(node.Left, element)
	} else {
		// 与右子节点的元素进行比较
		bst.add(node.Rigth, element)
	}

}