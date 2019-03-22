package ArrayQueue

import (
	"bytes"
	"fmt"
	"themoonstone/DataStruct/cap3-stack-queue/c03-queue/array"
)

type ArrayQueue struct {
	Array  array.Array
}

func (aq *ArrayQueue) Constructor() {
	aq.Array.InitDefaultArray()
}

func (aq *ArrayQueue) ConstructorArg(cap int) {
	aq.Array.InitArray(cap)
}

func (aq *ArrayQueue) Enqueue(element interface{})  {
	aq.Array.AddLast(element)
}

func (aq *ArrayQueue) Dequeue() interface{}{
	return aq.Array.RemoveHead()
}

func (aq *ArrayQueue) GetFront( ) interface{} {
	return aq.Array.GetHead()
}

func (aq *ArrayQueue) GetSize() int {
	return aq.Array.GetSize()
}

func (aq *ArrayQueue) IsEmpty() bool{
	return aq.Array.IsEmpty()
}

func (aq *ArrayQueue) GetCap() int {
	return aq.Array.GetCapacity()
}

// 格式化输出
func (aq *ArrayQueue) String() string {
	var result bytes.Buffer
	var s []byte
	s = append(s,[]byte(fmt.Sprintf("\tthe cap is [%d], the size is [%d]\n", aq.GetCap(), aq.GetSize()))...)
	s = append(s,[]byte("Queue : ")...)
	s = append(s,[]byte("head [")...)
	for i := 0; i < aq.GetSize(); i++{
		s = append(s,[]byte(fmt.Sprintf("%v",aq.Array.Get(i)))...)
		if i != aq.GetSize() - 1 {
			s = append(s, []byte(",")...)
		}
	}

	s = append(s,[]byte("] tail")...)
	result.Write(s)
	return result.String()
}