package basic

type BasicTree struct {
	Root *Node
	size int
}

type Node struct {
	Left    *Node
	Rigth   *Node
	Element interface{}
}

// 初始化
func (tree *BasicTree) Constructor(e interface{}) {
	tree.Root.Element = e
	tree.Root.Left = nil
	tree.Root.Rigth = nil
}

// 大小
func (tree *BasicTree) Size() int {
	return tree.size
}

// 是否为空
func (tree *BasicTree) IsEmpty() bool {
	return tree.size == 0
}
