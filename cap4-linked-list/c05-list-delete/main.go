package main

import (
	"fmt"
	"themoonstone/DataStruct/cap4-linked-list/c05-list-delete/list"
)

func main() {
	var List list.LinkedList
	List.Constructor()
	// 插入
	for i := 0; i < 5; i++ {
		List.InsertHead(i)
	}
	fmt.Println(List.String())
	// 向指定位置插入
	List.Insert(3,100)
	fmt.Println(List.String())

	// 查询指定节点
	fmt.Println(List.Get(5))
	// 查询指定元素是否包含在链表中
	fmt.Println(List.Contains(200))
	// 修改节点
	List.Update(2, 1000)
	fmt.Println(List.String())
	// 删除节点
	List.Delete(2)
	fmt.Println(List.String())
	List.DeleteHead()
	fmt.Println(List.String())
	List.DeleteTail()
	fmt.Println(List.DeleteTail())
}
