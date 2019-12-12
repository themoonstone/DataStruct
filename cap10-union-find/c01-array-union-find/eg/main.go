package main

import (
	"fmt"
	c01_array_union_find "github.com/themoonstone/DataStruct/cap10-union-find/c01-array-union-find"
)

func main() {
	auf := c01_array_union_find.Construct(10)
	fmt.Println(auf.IsConnected(3,5))
	auf.Merge(3, 5)
	fmt.Println(auf.IsConnected(3,5))
}
