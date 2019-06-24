package set

// 集合接口定义
type Set interface {
	// 添加元素
	Add(e interface{})
	// 删除元素
	Remove(e interface{})
	// 包含元素
	Contains(e interface{}) bool
	// 获取集合大小
	GetSize() int
	// 集合非空判断
	IsEmpty() bool
}
