package main

import (
	"fmt"
	"themoonstone/DataStruct/cap3-stack-queue/c05-loop-basic-queue/ArrayQueue"
	"themoonstone/DataStruct/cap3-stack-queue/c05-loop-basic-queue/LoopQueue"
	"themoonstone/DataStruct/cap3-stack-queue/c05-loop-basic-queue/queue"
	"time"
)

func QueueOp(queue queue.Queue,opCount int) int64{
	start := time.Now().UnixNano()
	for i := 0; i < opCount; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}
	end := time.Now().UnixNano()

	return (end - start)
}

func main() {
	var basicQueue ArrayQueue.ArrayQueue
	basicQueue.Constructor()
	// 注意，此处必需添加引用，否则会报错:method has pointer receiver
	// 这是因为golang中的method机制导致的，值类型的对象只有（t T) 结构的方法，虽然值类型的对象也可以调用(t *T) 方法，
	// 但这实际上是Golang编译器自动转化成了&t的形式来调用方法，并不是表明值类型的对象拥有该方法。
	// 这就意味着指针类型的receiver 方法实现接口时，只有指针类型的对象实现了该接口
	// 也就是说，对于我们当前的实例，只有&basicQueue实现了queue接口，而basicQueue没有实现该接口，所以需要添加引用

	fmt.Printf("basicQueue:%v\n", QueueOp(&basicQueue, 100000))
	var loopQueue LoopQueue.LoopQueue
	loopQueue.Constructor()
	fmt.Printf("loopQueue:%v\n", QueueOp(&loopQueue, 100000))
}
