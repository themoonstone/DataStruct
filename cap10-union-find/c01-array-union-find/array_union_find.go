package c01_array_union_find

type ArrayUnionFind struct {
	Size int
	UFS  []int
}

// 构造函数
func Construct(size int) *ArrayUnionFind {
	auf := &ArrayUnionFind{
		Size: size,
		UFS:  make([]int, size),
	}
	for i := 0; i < size; i++ {
		auf.UFS[i] = i 	// 初始化
	}
	return auf
}
// 查询
func (auf *ArrayUnionFind) IsConnected(p, q int) bool {
	return auf.find(p) == auf.find(q)
}
// 合并
func (auf *ArrayUnionFind) Merge(p, q int) {
	pID := auf.find(p)
	qID := auf.find(q)
	if pID == qID {
		return
	}
	for i := 0; i < auf.Size; i++ {
		if auf.UFS[i] == pID {
			auf.UFS[i] = qID
		}
	}
}
// 查找
func (auf *ArrayUnionFind) find(i int) int {
	if i < 0 || i > auf.Size - 1 {
		panic("index out of range")
	}
	return auf.UFS[i]
}
