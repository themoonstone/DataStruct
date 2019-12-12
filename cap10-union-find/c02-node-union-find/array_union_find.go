package c02_node_union_find

type ArrayUnionFind struct {
	Size int
	Parent  []int
}

// 构造函数
func Construct(size int) *ArrayUnionFind {
	auf := &ArrayUnionFind{
		Size: size,
		Parent:  make([]int, size),
	}
	for i := 0; i < size; i++ {
		auf.Parent[i] = i 	// 初始化、每个节点的根节点初始化都是自己
	}
	return auf
}
// 查询
func (auf *ArrayUnionFind) IsConnected(p, q int) bool {
	return auf.find(p) == auf.find(q)
}
// 合并
func (auf *ArrayUnionFind) Merge(p, q int) {
	pRoot := auf.find(p)
	qRoot := auf.find(q)
	if pRoot == qRoot {
		return
	}
	auf.Parent[pRoot] = qRoot
}
// 查找,只要i没有指向自己、就不断查找、走到找到根节点
func (auf *ArrayUnionFind) find(i int) int {
	if i < 0 || i > auf.Size - 1 {
		panic("index out of range")
	}
	for i != auf.Parent[i] {
		i = auf.Parent[i] // 一直向上查找
	}
	return i
}
