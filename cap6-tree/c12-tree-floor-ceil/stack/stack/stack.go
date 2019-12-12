package stack

// 实现一个栈的接口
type Stack interface {
	// 入栈
	Push(element interface{})
	// 出栈
	Pop() interface{}
	// 查看栈顶元素
	Peek() interface{}
	// 获取栈的大小
	GetSize() int
	// 检查栈是否为空
	IsEmpty() bool
}
