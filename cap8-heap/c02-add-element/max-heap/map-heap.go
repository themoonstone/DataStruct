package max_heap

import (
	"themoonstone/DataStruct/cap6-tree/utils/interfaces"
	"themoonstone/DataStruct/cap8-heap/c02-add-element/array"
)

// 使用数组实现堆
type MaxHeap struct {
	array.Array
}

// 无参构造函数
func (heap *MaxHeap) Construct() {
	heap.InitDefaultArray()
}
// 有参构造函数
func (heap *MaxHeap) ConstructWithCap(cap int)  {
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
	return 2 * index + 1
}
// 获取右子节点索引
func (heap *MaxHeap) RigthChild(index int) int {
	return 2 * (index + 1)
}

// 添加节点
func (heap *MaxHeap) Add(element interface{}) {
	heap.AddLast(element)
	// 上浮
	heap.shiftUp(heap.GetSize() - 1)
}

// 堆节点的上浮
func (heap *MaxHeap) shiftUp(index int)  {
	// 对当前插入的元素与其父节点元素大小进行判断,如果大于父节点，则交换位置继续判断
	for index != 0 && interfaces.Compare(heap.Get(index), heap.Get(heap.Parent(index))) == 1{
		// 交换
		heap.Swap(index, heap.Parent(index))
		// 更新index
		index = heap.Parent(index)
	}
}