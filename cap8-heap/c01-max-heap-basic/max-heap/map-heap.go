package max_heap

import "themoonstone/DataStruct/cap8-heap/c01-max-heap-basic/array"

// 使用数组实现堆
type MaxHeap struct {
	array.Array
}

// 无参构造函数
func (heap *MaxHeap) Construct() {
	heap.InitDefaultArray()
}

// 有参构造函数
func (heap *MaxHeap) ConstructWithCap(cap int) {
	heap.InitArray(cap)
}

// 大小
func (heap *MaxHeap) Size() int {
	return heap.GetSize()
}

// 非空判断
func (heap *MaxHeap) IsEmpty() bool {
	return heap.GetSize() == 0
}

// 获取父节点索引
func (heap *MaxHeap) Parent(index int) int {
	// 如果本身就是根节点，报错
	if index == 0 {
		panic("index-0 have not parent index")
	}
	return (index - 1) / 2
}

// 获取左子节点索引
func (heap *MaxHeap) LeftChild(index int) int {
	return 2*index + 1
}

// 获取右子节点索引
func (heap *MaxHeap) RigthChild(index int) int {
	return 2 * (index + 1)
}
