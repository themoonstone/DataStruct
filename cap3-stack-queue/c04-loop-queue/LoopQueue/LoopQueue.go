package LoopQueue
// 循环队列
type LoopQueue struct {
	Front		int
	Tail		int
	Size		int
	data		[]interface{}
}

// 队列初始化
func (lq *LoopQueue) Constructor() {
	lq.ConstructorArg(10)
}

// 队列初始化(带参)
func (lq *LoopQueue) ConstructorArg(cap int) {
	lq.data = make([]interface{}, cap)
}

// 获取队列长度
func (lq *LoopQueue) GetSize() int {
	return lq.Size
}
// 获取容积
func (lq *LoopQueue) GetCap() int {
	return len(lq.data) - 1// 注意，这里有意减1，是为了为循环队列预留一个空间，也就是说整个底层数组中，在还剩下一个空间的时候我们就认为队列已满，有意浪费一个空间
	// 需要注意这里的原因，根据readme.md中，front==tail表示队列为空，如果不预留一个空间，最终队列满也是front==tail，这样会表示两种情况，所以有意进行了避免
}

// 判断队列是否为空
func (lq *LoopQueue) IsEmpty() bool {
	return lq.Front == lq.Tail
}