package mapImplement

import (
	"github.com/labstack/gommon/log"
	"themoonstone/DataStruct/cap7-set/c05-bstree-map/utils/interfaces"
)

// 基于二分搜索树实现map

type BasicTreeMap struct {
	Root *Node
	size int
}

type Node struct {
	Left  *Node
	Right *Node
	Key   interface{} // key
	Value interface{} // value
}

// 初始化
func (tree_map *BasicTreeMap) Constructor() {
	tree_map.size = 0
	tree_map.Root = nil
}

// 生成一个新的节点
func genNode(key, value interface{}) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Key:   key,
		Value: value,
	}
}

func (tree_map *BasicTreeMap) get(node *Node, key interface{}) *Node {
	// key值判断
	// 递归终止条件， node为空或者找到key所属节点
	if node == nil {
		return nil
	}
	if interfaces.Compare(node.Key, key) == 0 {
		return node
	} else if interfaces.Compare(node.Key, key) == 1 {
		return tree_map.get(node.Left, key)
	} else {
		return tree_map.get(node.Right, key)
	}
}

// 插入
func (tree_map *BasicTreeMap) Add(key interface{}, value interface{}) {
	tree_map.Root = tree_map.add(tree_map.Root, key, value)
}

// 插入递归实现
// 返回新的二分搜索树的根节点
func (tree_map *BasicTreeMap) add(node *Node, key, value interface{}) *Node {
	if node == nil {
		tree_map.size++
		return genNode(key, value)
	}
	if interfaces.Compare(node.Key, key) == 1 {
		node.Left = tree_map.add(node.Left, key, value)
	} else if interfaces.Compare(node.Key, key) == -1 {
		node.Right = tree_map.add(node.Right, key, value)
	} else {
		node.Key = key
		node.Value = value
	}
	return node
}

// 查找最小节点
func (tree_map *BasicTreeMap) GetMinNode() interface{} {
	// 如果节点数量为0，抛出异常
	if tree_map.Size() == 0 {
		panic("the tree is empty!")
	}
	return tree_map.minNode(tree_map.Root).Key
}

// 查找最小节点递归实现
func (tree_map *BasicTreeMap) minNode(node *Node) *Node {
	// 递归终止条件
	// left已经为空了
	if node.Left == nil {
		return node
	}
	// 缩减规模
	return tree_map.minNode(node.Left)
}

// 删除二分搜索树最小节点
func (tree_map *BasicTreeMap) RemoveMin() interface{} {
	// 找到最小节点
	min := tree_map.GetMinNode()
	tree_map.Root = tree_map.removeMin(tree_map.Root)
	return min
}

// 删除以node为根节点的二分搜索树最小节点(递归实现)
// 返回删除节点后的新的二分搜索树的根
func (tree_map *BasicTreeMap) removeMin(node *Node) *Node {
	// 删除该节点
	// 删除的时候应该分两种情况
	// 1. 待删除的节点是叶子节点，直接删除
	// 2. 待删除的节点不是叶子节点，它还有右子树
	// 递归终止条件 node.left=nil,说明该节点即是最小节点，需要删除该节点
	if node.Left == nil {
		rightNode := node.Right
		node.Right = nil
		tree_map.size--
		return rightNode
	}
	// 递归判断，一直到左子树为空
	// 删除node的左子树对应的最小值，二分搜索树本身还是以node为根，所以返回的是node
	node.Left = tree_map.removeMin(node.Left)
	return node
}

// 查找最大节点
func (tree_map *BasicTreeMap) GetMaxNode() interface{} {

	// 如果节点数量为0，抛出异常
	if tree_map.Size() == 0 {
		panic("the tree is empty!")
	}
	return tree_map.maxNode(tree_map.Root).Key
}

// 查找最大节点递归实现
func (tree_map *BasicTreeMap) maxNode(node *Node) *Node {
	// 递归终止条件
	// rigth已经为空了
	if node.Right == nil {
		return node
	}
	// 缩减规模
	return tree_map.maxNode(node.Right)
}

// 删除最大节点
// 返回删除的最大值
func (tree_map *BasicTreeMap) RemoveMax() interface{} {
	max := tree_map.GetMaxNode()
	// 删除指定根节点的二分搜索树的最大节点
	tree_map.Root = tree_map.removeMax(tree_map.Root)
	return max
}

// 删除二分搜索树最大节点递归实现
// 删除指定根节点的二分搜索树的最大节点
// 返回删除节点之后的新的二分搜索树的根节点
// node:表示要删除的二分搜索树的根节点
func (tree_map *BasicTreeMap) removeMax(node *Node) *Node {
	// 递归终止条件
	if node.Right == nil {
		leftNode := node.Left
		node.Left = nil
		tree_map.size--
		return leftNode
	}
	// 缩减规模
	node.Right = tree_map.removeMax(node.Right)
	return node
}

func (tree_map *BasicTreeMap) Remove(key interface{}) interface{} {
	node := tree_map.get(tree_map.Root, key)
	if nil != node {
		tree_map.Root = tree_map.remove(tree_map.Root, key)
		return node.Value
	}
	return nil
}

// 删除递归实现
func (tree_map *BasicTreeMap) remove(node *Node, key interface{}) *Node {
	// 递归终止条件，node为空或者删除完成
	if node == nil {
		return nil
	}
	if interfaces.Compare(node.Key, key) == 1 {
		node.Left = tree_map.remove(node.Left, key)
		return node
	} else if interfaces.Compare(node.Key, key) == -1 {
		node.Right = tree_map.remove(node.Right, key)
		return node
	} else {
		// 删除
		// node只有左子树
		if node.Right == nil {
			leftNode := node.Left
			// go语言有自动回收，理论上可以直接返回node.Left
			node.Left = nil
			// 总节点数量减一
			tree_map.size--
			return leftNode
		} else if node.Left == nil {
			// node只有右子树
			rightNode := node.Right
			node.Right = nil
			tree_map.size--
			return rightNode
		} else {
			// node既有左子树，又有右子树
			// 3. 如果待删除节点既有左子树又有右子树
			// 3.1 找到后继节点进行缓存
			successor := tree_map.minNode(node.Right)
			// 3.2 删除后继节点
			successor.Right = tree_map.removeMin(node.Right)
			// 3.3 以保存的successor后继节点顶替当前节点
			successor.Left = node.Left
			// 3.4 删除当前节点
			node.Right = nil
			node.Left = nil
			// 返回successor, 以后继节点为删除指定节点之后的新的二分搜索树的根节点
			return successor
		}
	}
}

func (tree_map *BasicTreeMap) Set(key, value interface{}) {
	node := tree_map.get(tree_map.Root, key)
	if node == nil {
		log.Panicf("key [%v] is not exists\n", key)
	}
	node.Value = value
}

func (tree_map *BasicTreeMap) Get(key interface{}) interface{} {
	node := tree_map.get(tree_map.Root, key)
	if node == nil {
		return nil
	} else {
		return node.Value
	}
}

func (tree_map *BasicTreeMap) Contains(key interface{}) bool {
	if tree_map.get(tree_map.Root, key) != nil {
		return true
	}
	return false
}

func (tree_map *BasicTreeMap) IsEmpty() bool {
	return tree_map.size == 0
}

func (tree_map *BasicTreeMap) Size() int {
	return tree_map.size
}
