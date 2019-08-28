package c01_line_segment_tree_basic

type LineSegmentTree struct {
	//
	// 数组副本(对于区间的每一个元素,用户有可能需要通过线段树进行获取,所以这里设置一个数组的副本)
	Data []interface{}
	// 数组长度
	size int
}

// 构造函数(用户传进来的应该是整个区间的范围)
func (segmentTree *LineSegmentTree)Constructor(data []interface{})  {
	segmentTree.Data = make([]interface{},len(data))
	segmentTree.Data = data
	segmentTree.size = len(data)
}

// 返回长度
func (segmentTree *LineSegmentTree) GetSize() int {
	return segmentTree.size
}

// 获取指定下标的元素
func (segmentTree *LineSegmentTree) Get(index int) interface{} {
	if index >= segmentTree.size {
		panic("index out of boundary")
	}
	return segmentTree.Data[index]
}

// 返回一个完全二叉树的数组表示中,一个索引所表示的左孩子的节点的索引
func (segmentTree *LineSegmentTree) leftChild(index int) int {
	return 2 * index + 1
}

// 返回一个完全二叉树的数组表示中,一个索引所表示的右孩子的节点的索引
func (segmentTree *LineSegmentTree) rightChild(index int) int {
	return 2 * index + 2
}