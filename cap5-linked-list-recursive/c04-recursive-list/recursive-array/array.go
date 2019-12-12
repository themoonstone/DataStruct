package recursive_array

import "fmt"

type RecursiveArray struct {
}

func (rs *RecursiveArray) Sum(array []int) int {

	return rs.sum(0, array)
}

// leftValue:左边界
// array:不断变小的数组
func (rs *RecursiveArray) sum(leftValue int, array []int) int {
	if leftValue == len(array)-1 {
		return array[leftValue]
	}
	fmt.Println("length : ", len(array), "leftValue : ", leftValue)
	// 从计算leftValue到n(数组末尾)的数字的和到变成从leftValue+1到n的数字的和，不断的缩减问题的规模
	return array[leftValue] + rs.sum(leftValue+1, array)
}
