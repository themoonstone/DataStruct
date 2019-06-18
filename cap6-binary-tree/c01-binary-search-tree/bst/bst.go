package bst

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