package LoopQueue

import (
	"bytes"
	"fmt"
)

// 循环队列
type LoopQueue struct {
	Front int
	Tail  int
	Size  int
	data  []interface{}
}

// 队列初始化
func (lq *LoopQueue) Constructor() {
	lq.ConstructorArg(10)
}

// 队列初始化(带参)
// 注意，用户传过来cap是希望我们最大可以承载cap个元素
// 但循环队列中，我们有意识的浪费了一个空间，所以应该是cap+1
func (lq *LoopQueue) ConstructorArg(cap int) {
	lq.data = make([]interface{}, cap+1)
}

// 获取队列长度
func (lq *LoopQueue) GetSize() int {
	return lq.Size
}

// 获取容积
func (lq *LoopQueue) GetCap() int {
	return len(lq.data) - 1 // 注意，这里有意减1，是为了为循环队列预留一个空间，也就是说整个底层数组中，在还剩下一个空间的时候我们就认为队列已满，有意浪费一个空间
	// 需要注意这里的原因，根据readme.md中，front==tail表示队列为空，如果不预留一个空间，最终队列满也是front==tail，这样会表示两种情况，所以有意进行了避免
}

// 判断队列是否为空
func (lq *LoopQueue) IsEmpty() bool {
	return lq.Front == lq.Tail
}

// 判断队列是否已满
func (lq *LoopQueue) IsFull() bool {
	return (lq.Tail+1)%len(lq.data) == lq.Front
}

// 入队
func (lq *LoopQueue) Enqueue(element interface{}) {
	// 判断队列是否已满
	if lq.IsFull() {
		// 扩容,取现在的循环队列最多可以存放的元素个数*2
		lq.resise(2 * lq.GetCap())
	}
	lq.data[lq.Tail] = element
	// 这里因为是一个循环队列，所以不能直接自加
	lq.Tail = (lq.Tail + 1) % len(lq.data)
	lq.Size++
}

// 出队
func (lq *LoopQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		panic("Queue is Empty")
	}
	// 获取队头元素
	element := lq.data[lq.Front]
	lq.data[lq.Front] = nil
	lq.Front = (lq.Front + 1) % len(lq.data)
	lq.Size--
	if lq.Size < lq.GetCap()/4 {
		lq.resise(lq.GetCap() / 2)
	}
	return element
}

// 获取队头元素
func (lq *LoopQueue) GetFront() interface{} {
	if lq.IsEmpty() {
		panic("Queue is Empty")
	}
	return lq.data[lq.Front]
}

// 队列容积重新分配
func (lq *LoopQueue) resise(newCap int) {
	newData := make([]interface{}, newCap+1)
	for i := 0; i < lq.Size; i++ {
		newData[i] = lq.data[(i+lq.Front)%len(lq.data)]
	}

	lq.data = newData
	lq.Front = 0
	lq.Tail = lq.Size
}

// 格式化输出
func (lq *LoopQueue) String() string {
	var result bytes.Buffer
	var s []byte
	s = append(s, []byte(fmt.Sprintf("\tthe cap is [%d], the size is [%d]\n", lq.GetCap(), lq.GetSize()))...)
	s = append(s, []byte("Queue : ")...)
	s = append(s, []byte("front [")...)
	// 遍历的索引i不能指向tail
	for i := lq.Front; i != lq.Tail; i = (i + 1) % len(lq.data) {
		s = append(s, []byte(fmt.Sprintf("%v", lq.data[i]))...)

		if (i+1)%len(lq.data) != lq.Tail {
			// 当前索引不是最后一个元素，因为是循环队列，所以和数组队列有很大的不同
			s = append(s, []byte(",")...)
		}
	}

	s = append(s, []byte("] tail")...)
	result.Write(s)
	return result.String()
}
