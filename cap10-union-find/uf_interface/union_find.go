package uf_interface

// 并查集接口
type UnionFind interface {
	IsConnected(p, q int) bool // 查询
	Merge(p, q int)            // 合并
}
