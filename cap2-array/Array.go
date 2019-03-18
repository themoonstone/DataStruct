package main

import "fmt"

type Array struct {
	data    []int		// 设置数组变量
	cap		int			// 设置数组容量
}
// 初始化
func (a *Array) InitArray(capacity int) {
	a.cap = capacity
	a.data = make([]int, 0, a.cap)
}

func (a *Array) InitDefaultArray() {
	a.cap = 10
	a.data = make([]int, 0, a.cap)
}
// 获取数组容量
func  (a *Array) GetCapacity() int {
	return a.cap
}
// 获取数组元素个数
func (a *Array) GetSize() int {
	return len(a.data)
}
// 获取数组是否为空
func (a *Array) IsEmpty() bool {
	return len(a.data) == 0
}
func main() {
	var a Array
	a.InitDefaultArray()
	fmt.Println(a.GetSize())
}
