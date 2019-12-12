package queue

// 队列接口
type Queue interface {
	// 出队
	Enqueue(element interface{})
	// 出队
	Dequeue() interface{}
	// 获取队头元素
	GetFront() interface{}
	// 获取队列长度
	GetSize() int
	// 判断队列是否为空
	IsEmpty() bool
}
