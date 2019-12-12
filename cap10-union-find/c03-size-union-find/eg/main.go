package main

import (
	"fmt"
	c03_size_union_find "github.com/themoonstone/DataStruct/cap10-union-find/c03-size-union-find"
)

func main() {
	auf := c03_size_union_find.Construct(10)
	fmt.Println(auf.IsConnected(3, 5))
	auf.Merge(3, 5)
	fmt.Println(auf.IsConnected(3, 5))
	auf.Merge(3, 9)
	fmt.Println(auf.IsConnected(3, 9))
	fmt.Println(auf)
}
