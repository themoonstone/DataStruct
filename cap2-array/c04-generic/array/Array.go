package array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{} // 设置数组变量
	size int           // 设置数组长度
}

// 初始化
func (a *Array) InitArray(capacity int) {
	a.data = make([]interface{}, capacity)
}

func (a *Array) InitDefaultArray() {
	a.data = make([]interface{}, 10)
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
func (a *Array) AddLast(val interface{}) {
	a.Insert(a.size, val)
}

// 向指定位置插入一个新的元素
func (a *Array) Insert(index int, val interface{}) {
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
func (a *Array) AddFirst(val interface{}) {
	a.Insert(0, val)
}

// 实现获取指定位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("index out of range")
	}
	return a.data[index]
}

// 实现修改指定位置的元素
func (a *Array) Set(index int, value interface{}) {
	if index < 0 || index >= a.size {
		panic("index out of range")
	}
	a.data[index] = value
}

// 实现包含函数
func (a *Array) Contain(value interface{}) bool {
	for i := 0; i < a.size; i++ {
		if value == a.data[i] {
			return true
		}
	}
	return false
}

// 实现查找函数
// 查找指定元素的索引，如果没有找到返回-1
func (a *Array) Find(value interface{}) int {
	for i := 0; i < a.size; i++ {
		if value == a.data[i] {
			return i
		}
	}
	return -1
}

// 实现删除函数
// 返回被删除元素
func (a *Array) Remove(index int) (value interface{}) {
	if index < 0 || index >= a.size {
		panic("index out of range")
	}
	value = a.data[index]
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	return
}

// 删除头部元素
func (a *Array) RemoveHead() (value interface{}) {
	return a.Remove(0)
}

// 删除尾部元素
func (a *Array) RemoveTail() (value interface{}) {
	return a.Remove(a.size - 1)
}

// 实现自定义数组的格式化输出
func (a *Array) String() string {
	var result bytes.Buffer
	var s []byte
	s = append(s, []byte("[")...)
	for i := 0; i < a.size; i++ {
		s = append(s, []byte(fmt.Sprintf("%v", a.data[i]))...)
		if i != a.size-1 {
			s = append(s, []byte(",")...)
		}
	}

	s = append(s, []byte("]")...)
	result.Write(s)
	return result.String()
}
