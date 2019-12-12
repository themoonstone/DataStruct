package _map

// map类型接口实现

type Map interface {
	Add(key interface{}, value interface{})
	Remove(key interface{}) interface{}
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Contains(key interface{}) bool
	IsEmpty() bool
	Size() int
}
