package SetImplement

import "themoonstone/DataStruct/cap7-set/c02-link-list-set/linked-list"

type LinkedListSet struct {
	linked_list.LinkedList
}

// set 接口的实现
func (ls *LinkedListSet) Add(e interface{}) {
	// 根据集合的特点，需要去重
	if !ls.Contains(e) {
		ls.LinkedList.InsertTail(e)
	}
}

func (ls *LinkedListSet) Remove(e interface{}) {
	ls.LinkedList.RemoveElement(e)
}

func (ls *LinkedListSet) IsEmpty() bool {
	return ls.LinkedList.IsEmpty()
}

func (ls *LinkedListSet) Contains(e interface{}) bool {
	return ls.LinkedList.Contains(e)
}

func (ls *LinkedListSet) GetSize() int {
	return ls.LinkedList.GetSize()
}
