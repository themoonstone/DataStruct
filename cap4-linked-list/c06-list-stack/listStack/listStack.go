package listStack

import (
	"fmt"
	"themoonstone/DataStruct/cap4-linked-list/c06-list-stack/list"
)

// 实现一个栈的接口
//type Stack interface {
//	// 入栈
//	Push(element interface{})
//	// 出栈
//	Pop() interface{}
//	// 查看栈顶元素
//	Peek() interface{}
//	// 获取栈的大小
//	GetSize() int
//	// 检查栈是否为空
//	IsEmpty() bool
//}

// 代表栈基本结构
type ListStack struct {
	linkList list.LinkedList
}

// 构造函数
func (ls *ListStack) Constructor() {
	ls.linkList.Constructor()
}

func (ls *ListStack) Push(element interface{}) {
	ls.linkList.InsertHead(element)
}

func (ls *ListStack) Pop() interface{} {
	return ls.linkList.DeleteHead()
}
func (ls *ListStack) Peek() interface{} {
	return ls.linkList.GetFirst()
}
func (ls *ListStack) GetSize() int {
	return ls.linkList.GetSize()
}
func (ls *ListStack) IsEmpty() bool {
	return ls.linkList.IsEmpty()
}

// 实现格式化输出
func (ls *ListStack) String() string {
	var res []interface{}
	for cur := ls.linkList.DummyHead.Next; cur != nil; cur = cur.Next {
		res = append(res, cur.Element)
		res = append(res, "->")
	}
	res = append(res, "NULL")
	return fmt.Sprintf("Stack:top %v\n", res)
}
