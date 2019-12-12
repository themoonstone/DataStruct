package main

import (
	"fmt"
	c02_node_union_find "github.com/themoonstone/DataStruct/cap10-union-find/c02-node-union-find"
)

func main() {
	auf := c02_node_union_find.Construct(10)
	fmt.Println(auf.IsConnected(3, 5))
	auf.Merge(3, 5)
	fmt.Println(auf.IsConnected(3, 5))
	auf.Merge(3, 9)
	fmt.Println(auf.IsConnected(3, 9))
	fmt.Println(auf)
}
