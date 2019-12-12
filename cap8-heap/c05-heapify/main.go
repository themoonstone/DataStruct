package main

import (
	"fmt"
	"log"
	"math/rand"
	"themoonstone/DataStruct/cap8-heap/c05-heapify/max-heap"
	"themoonstone/DataStruct/cap8-heap/utils/interfaces"
	"time"
)

var maxN = 2000000

func main() {
	// 插入1000000个随机数

	//var heap max_heap.MaxHeap
	// 将这1000000个数存入切片
	var slice []interface{} = make([]interface{}, maxN)
	for i := 0; i < maxN; i++ {
		slice[i] = rand.Int63n(int64(maxN))
	}
	fmt.Printf("heapify time: %v\n", testHeap(slice, true))

	fmt.Printf("normal time: %v\n", testHeap(slice, false))
}

func testHeap(arr []interface{}, isHeapify bool) float64 {
	start := time.Now().UnixNano()
	var heap max_heap.MaxHeap
	if isHeapify {

		heap.ConstructWithArray(arr)
	} else {
		heap.Construct()
		for i := 0; i < len(arr); i++ {
			heap.Add(arr[i])
		}
	}

	// 将这1000000个数存入切片
	var slice []interface{} = make([]interface{}, maxN)
	for i := 0; i < maxN; i++ {
		slice[i] = heap.Remove()
	}
	// 遍历切片，进行比较，如果出现前一个数比后一个数小，则说明堆的实现有问题
	for i := 0; i < maxN-1; i++ {
		if interfaces.Compare(slice[i], slice[i+1]) == -1 {
			log.Panicf("the struct is not correct!\n")
		}
	}
	fmt.Printf("test complete\n")
	end := time.Now().UnixNano()
	fmt.Printf("start : %d, end : %d\n", start, end)
	return float64((end - start)) / float64(1000000000.0)
}
