package main

import (
	"fmt"
	"log"
	"math/rand"
	"themoonstone/DataStruct/cap8-heap/c03-remove-element/max-heap"
)

func main() {
	// 插入1000000个随机数
	var maxN = 1000000
	var heap max_heap.MaxHeap
	heap.Construct()
	for i := 0; i < maxN; i++ {
		heap.Add(rand.Int63n(int64(maxN)))
	}
	// 将这1000000个数存入切片
	var slice []int64 = make([]int64, maxN)
	for i := 0; i < maxN; i++{
		slice[i] = heap.Remove().(int64)
	}
	// 遍历切片，进行比较，如果出现前一个数比后一个数小，则说明堆的实现有问题
	for i := 0; i < maxN - 1; i++ {
		if slice[i] < slice[i+1] {
			log.Panicf("the struct is not correct!\n")
		}
	}

	fmt.Printf("test completed!\n")
}
