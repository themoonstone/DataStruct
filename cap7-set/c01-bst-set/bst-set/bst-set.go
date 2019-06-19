package bst_set

import "themoonstone/DataStruct/cap7-set/c01-bst-set/bstree/basic"

// 基于二分搜索树的集合定义
type BSTSet struct {
	basic.BasicTree
}

func (bs *BSTSet) Construct() {
	bs.BasicTree.Constructor()
}

// set 接口的实现
func (bs *BSTSet) Add(e interface{})  {
	bs.BasicTree.Add(e)
}

func (bs *BSTSet) Remove(e interface{})  {
	bs.BasicTree.RemoveAnyNode(e)
}

func (bs *BSTSet) IsEmpty() bool {
	return bs.BasicTree.IsEmpty()
}

func (bs *BSTSet) Contains(e interface{}) bool {
	return bs.BasicTree.Contains(e)
}

func (bs *BSTSet) GetSize() int {
	return bs.BasicTree.Size()
}