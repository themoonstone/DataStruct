package arrayStack

import (
	"bytes"
	"fmt"
	"themoonstone/DataStruct/cap4-linked-list/c06-list-stack/array"
)

// 实现基于动态数组的栈
type ArrayStack struct {
	Array array.Array
}

func (as *ArrayStack) Push(element interface{}) {
	as.Array.AddLast(element)
}

func (as *ArrayStack) Pop() interface{} {
	return as.Array.RemoveTail()
}

func (as *ArrayStack) Peek() interface{} {
	return as.Array.GetLast()
}

func (as *ArrayStack) GetSize() int {
	return as.Array.GetSize()
}

// 获取容积
func (as *ArrayStack) GetCap() int {
	return as.Array.GetCapacity()
}

func (as *ArrayStack) IsEmpty() bool {

	return as.Array.IsEmpty()
}

// 添加一个构造函数
func (as *ArrayStack) Constructor() {
	as.Array.InitDefaultArray()
}

// 重写String，实现栈的格式化输出
func (as *ArrayStack) String() string {
	var result bytes.Buffer
	var s []byte
	s = append(s, []byte(fmt.Sprintf("\tthe cap is [%d], the size is [%d]\n", as.GetCap(), as.GetSize()))...)
	s = append(s, []byte("Stack : [")...)
	for i := 0; i < as.GetSize(); i++ {
		s = append(s, []byte(fmt.Sprintf("%v", as.Array.Get(i)))...)
		if i != as.GetSize()-1 {
			s = append(s, []byte(",")...)
		}
	}

	s = append(s, []byte("] top")...)
	result.Write(s)
	return result.String()
}
