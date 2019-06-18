package main

import (
	"fmt"
	"themoonstone/DataStruct/cap6-tree/c10-tree-del-MNode/basic"
)

func main() {
	tree := basic.BasicTree{}
	tree.Constructor()
	var data []int = []int{5,3,6,8,4,2}
	for i := 0; i < len(data); i++ {
		tree.Add(data[i])
	}
	//fmt.Println(tree.String())
	// 前序遍历
	tree.FrontIter()
	tree.FrontIterNR()

	// 5 3 6 2 4 8
	tree.LayerIter()
	// 中序遍历
	//tree.AfterIter()
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
	tree.LayerIter()
}
