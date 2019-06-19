package main

import (
	"fmt"
	"themoonstone/DataStruct/cap6-tree/c12-tree-floor-ceil/basic"
)

func main() {
	tree := basic.BasicTree{}
	tree.Constructor()
	var data []int = []int{7,3,8,2,4,9}
	for i := 0; i < len(data); i++ {
		tree.Add(data[i])
	}
	/*
	//fmt.Println(tree.String())
	// 前序遍历
	tree.FrontIter()
	tree.FrontIterNR()

	// 5 3 6 2 4 8
	tree.LayerIter()
	// 后序遍历
	tree.AfterIter()
	// 查找最小节点
	fmt.Printf("\nmin : %v\n", tree.GetMinNode())

	// 删除最小节点
	tree.RemoveMin()
	fmt.Printf("\nmin : %v\n", tree.GetMinNode())
	// 5 3 6 2 4 8
	tree.LayerIter()
	// 查找最大节点
	fmt.Printf("\nmax : %v\n", tree.GetMaxNode())
	// 删除最大节点
	tree.RemoveMax()
	// 查找最大节点
	fmt.Printf("\nmax : %v\n", tree.GetMaxNode())
	tree.FrontIter()
	// 删除指定元素的节点
	tree.RemoveAnyNode(3)
	tree.FrontIter()
	*/
	// 查找指定元素的floor节点
	fmt.Printf("floor : %v\n", tree.Floor(4))

	// 查找指定元素的ceil节点
	//fmt.Printf("ceil : %v\n", tree.Ceil(7))
	fmt.Printf("ceil : %v\n", tree.Ceil(4))
	//fmt.Printf("ceil : %v\n", tree.Ceil(1))
}
