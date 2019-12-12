package max_heap

import (
	"themoonstone/DataStruct/cap8-heap/c03-remove-element/array"
	"themoonstone/DataStruct/cap8-heap/utils/interfaces"
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

// 添加节点
func (heap *MaxHeap) Add(element interface{}) {
	heap.AddLast(element)
	// 上浮
	heap.shiftUp(heap.GetSize() - 1)
}

// 堆节点的上浮
func (heap *MaxHeap) shiftUp(index int) {
	// 对当前插入的元素与其父节点元素大小进行判断,如果大于父节点，则交换位置继续判断

	for index != 0 && interfaces.Compare(heap.Get(index), heap.Get(heap.Parent(index))) == 1 {
		// 交换
		heap.Swap(index, heap.Parent(index))
		// 更新index
		index = heap.Parent(index)
	}
}

// 删除节点, 返回被删除的元素
func (heap *MaxHeap) Remove() interface{} {
	// 将头节点与尾部节点进行替换
	heap.Swap(0, heap.GetSize()-1)
	// 删除尾部节点
	value := heap.RemoveTail()
	// 下沉
	heap.shiftDown(0)
	return value
}

// 下沉
func (heap *MaxHeap) shiftDown(index int) {
	// 将待下沉的节点与左右子节点的最大值进行比较，如果小于子节点中的最大值，进行替换
	// 替换完成的条件，没有左子节点
	for heap.LeftChild(index) < heap.GetSize() {
		j := heap.LeftChild(index)
		// 如果当前堆有右子节点并且左子节点元素小于右子节点元素
		if heap.RigthChild(index) < heap.GetSize() && interfaces.Compare(heap.Get(j), heap.Get(j+1)) == -1 {
			j = heap.RigthChild(index)
		}
		// 判断最大值与index的值谁大
		if interfaces.Compare(heap.Get(index), heap.Get(j)) == 1 {
			// 如果Index的值大于子节点的最大值，不需要再下沉，直接break
			break
		}
		// 交换
		heap.Swap(index, j)
		index = j
	}
}
