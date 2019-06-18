package main

import (
	"themoonstone/DataStruct/cap6-tree/c08-tree-front-nr/basic"
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
	// 中序遍历
	//tree.AfterIter()
}
