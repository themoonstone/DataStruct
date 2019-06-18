package basic

import (
	"bytes"
	"fmt"
	"themoonstone/DataStruct/cap6-tree/c10-tree-del-MNode/queue/ArrayQueue"
	"themoonstone/DataStruct/cap6-tree/c10-tree-del-MNode/stack/arrayStack"
	"themoonstone/DataStruct/cap6-tree/utils/interfaces"
)

type BasicTree struct {
	Root *Node
	size	int
} 

type Node struct {
	Left 	*Node
	Right	*Node
	Element		interface{}
}

// 初始化
func (tree *BasicTree) Constructor()  {
	tree.size = 0
	tree.Root = nil
}
// 生成一个新的节点
func generate(element interface{}) *Node {
	return &Node{
		Left:nil,
		Right:nil,
		Element:element,
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
	if (interfaces.Compare(e, node.Element) == -1) {
		node.Left = tree.add(node.Left, e)
	} else if (interfaces.Compare(e, node.Element) == 1) {
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
		return tree.contains( node.Right, e)
	}
}

// 前序遍历
func (tree *BasicTree) FrontIter()  {
	tree.frontIter(tree.Root)
	fmt.Println()
}

// 前序遍历递归实现
func (tree *BasicTree) frontIter(node *Node) {
	// 递归终止条件：node is nil
	if node == nil {
		return
	}
	fmt.Printf(" %v ", node.Element)
	tree.frontIter(node.Left)
	tree.frontIter(node.Right)

}

// 前序遍历非递归实现
func (tree *BasicTree) FrontIterNR() {
	var stackTree arrayStack.ArrayStack
	stackTree.Array.InitDefaultArray()

	// 首先压入根节点
	stackTree.Push(tree.Root)
	// 对栈进行遍历操作
	for !stackTree.IsEmpty() {
		curNode := stackTree.Pop()
		//fmt.Printf("curNode : %v\n", reflect.TypeOf(curNode))
		// 需要添加断言
		if node, ok := curNode.(*Node);ok {
			fmt.Printf("%v ", node.Element)
			if node.Right != nil {
				stackTree.Push(node.Right)
			}
			if node.Left != nil {
				stackTree.Push(node.Left)
			}
		}
	}
	fmt.Println()
}

// 实现前序遍历的二分搜索树的打印输出
// 先打印根节点
// 再展示左子树
// 再展示右子树
func (tree *BasicTree) String() string {
	var buff bytes.Buffer
	tree.generateBSTString(tree.Root, 0, &buff)
	return buff.String()
}

// 中序遍历递归实现
func (tree *BasicTree)MidIter()  {
	tree.midIter(tree.Root)
}

// 中序遍历递归函数实现
func (tree *BasicTree) midIter(node *Node)  {
	// 递归终止条件
	if node == nil {
		return
	}

	// 缩减规模
	tree.midIter(node.Left)
	fmt.Printf("%v\n", node.Element)
	tree.midIter(node.Right)
}

// 后序遍历递归实现
func (tree *BasicTree)AfterIter()  {
	tree.afterIter(tree.Root)
}

// 后序遍历递归函数实现
func (tree *BasicTree) afterIter(node *Node)  {
	// 递归终止条件
	if node == nil {
		return
	}

	// 缩减规模
	tree.afterIter(node.Left)
	tree.afterIter(node.Right)
	fmt.Printf("%v\n", node.Element)
}

// 实现二分搜索树层序遍历(广度优先)
func (tree *BasicTree) LayerIter() {
	var queue ArrayQueue.ArrayQueue
	queue.Array.InitDefaultArray()
	queue.Enqueue(tree.Root)
	for !queue.IsEmpty() {
		curNode := queue.Dequeue()
		// 断言
		if node, ok := curNode.(*Node); ok {
			fmt.Printf("%v ", node.Element)
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}
}

// 前序遍历方式递归打印二叉树
func (tree *BasicTree) generateBSTString(node *Node, depth int, buff *bytes.Buffer) {
	if node == nil {
		// WriteString appends the contents of s to the buffer
		buff.WriteString(tree.generateDepthString(depth)+"nil\n")
		return
	}
	// --node.e
	buff.WriteString(fmt.Sprintf("%v%v\n", tree.generateDepthString(depth), node.Element))
	tree.generateBSTString(node.Left, depth+1,buff)
	tree.generateBSTString(node.Right, depth+1,buff)
	return
}
// 深度标识
func (tree *BasicTree) generateDepthString(depth int) string{
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

// 查找最小节点
func (tree *BasicTree) GetMinNode() interface{} {
	// 如果节点数量为0，抛出异常
	if tree.Size() == 0 {
		panic("the tree is empty!")
	}
	return tree.minNode(tree.Root).Element
}
// 查找最小节点递归实现
func (tree *BasicTree) minNode(node *Node) *Node {
	// 递归终止条件
	// left已经为空了
	if node.Left == nil {
		return node
	}
	// 缩减规模
	return tree.minNode(node.Left)
}

// 删除二分搜索树最小节点
func (tree *BasicTree) RemoveMin() interface{} {
	// 找到最小节点
	min := tree.GetMinNode()
	tree.Root = tree.removeMin(tree.Root)
	return min
}
// 删除以node为根节点的二分搜索树最小节点(递归实现)
// 返回删除节点后的新的二分搜索树的根
func (tree *BasicTree) removeMin(node *Node) *Node {
	// 删除该节点
	// 删除的时候应该分两种情况
	// 1. 待删除的节点是叶子节点，直接删除
	// 2. 待删除的节点不是叶子节点，它还有右子树
	// 递归终止条件 node.left=nil,说明该节点即是最小节点，需要删除该节点
	if node.Left == nil {
		rightNode := node.Right
		node.Right = nil
		tree.size--
		return rightNode
	}
	// 递归判断，一直到左子树为空
	// 删除node的左子树对应的最小值，二分搜索树本身还是以node为根，所以返回的是node
	node.Left = tree.removeMin(node.Left)
	return node
}

// 查找最大节点
func (tree *BasicTree) GetMaxNode() interface{} {

	// 如果节点数量为0，抛出异常
	if tree.Size() == 0 {
		panic("the tree is empty!")
	}
	return tree.maxNode(tree.Root).Element
}
// 查找最大节点递归实现
func (tree *BasicTree) maxNode(node *Node) *Node {
	// 递归终止条件
	// rigth已经为空了
	if node.Right == nil {
		return node
	}
	// 缩减规模
	return tree.maxNode(node.Right)
}

// 删除最大节点
// 返回删除的最大值
func (tree *BasicTree) RemoveMax() interface{} {
	max := tree.GetMaxNode()
	// 删除指定根节点的二分搜索树的最大节点
	tree.Root = tree.removeMax(tree.Root)
	return max
}

// 删除二分搜索树最大节点递归实现
// 删除指定根节点的二分搜索树的最大节点
// 返回删除节点之后的新的二分搜索树的根节点
// node:表示要删除的二分搜索树的根节点
func (tree *BasicTree) removeMax(node *Node) *Node {
	// 递归终止条件
	if node.Right == nil {
		leftNode := node.Left
		node.Left = nil
		tree.size--
		return leftNode
	}
	// 缩减规模
	node.Right = tree.removeMax(node.Right)
	return node
}