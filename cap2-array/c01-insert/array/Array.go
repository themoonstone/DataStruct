package array

type Array struct {
	data []int // 设置数组变量
	size int   // 设置数组长度
}

// 初始化
func (a *Array) InitArray(capacity int) {
	a.data = make([]int, capacity)
}

func (a *Array) InitDefaultArray() {
	a.data = make([]int, 10)
}

// 获取数组容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获取数组元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 获取数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 向数组尾部添加一个新的元素
func (a *Array) AddLast(val int) {
	a.Insert(a.size, val)
}

// 向指定位置插入一个新的元素
func (a *Array) Insert(index, val int) {
	if a.size == len(a.data) {
		panic("full")
	}
	// 判断index是否大于cap或者小于0
	if index < 0 && index > a.GetCapacity() {
		panic("error : index out of range")
	}
	// 将大于Index的位置的元素后移一位
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	// 将index的val插入
	a.data[index] = val
	a.size++
}

// 向数组头部插入一个新的元素
func (a *Array) AddFirst(val int) {
	a.Insert(0, val)
}
