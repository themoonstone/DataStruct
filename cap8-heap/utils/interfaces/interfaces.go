package interfaces


import "reflect"

// 实现 interface{}类型比较

func Compare(a, b interface{}) int {
	if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
		panic("unable to compare！")
	}
	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int){
			return 0
		} else {
			return -1
		}
	case int64:
		if a.(int64) > b.(int64) {
			return 1
		} else if a.(int64) == b.(int64){
			return 0
		} else {
			return -1
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) == b.(string){
			return 0
		} else {
			return -1
		}
	case float32:
		if a.(float32) > b.(float32) {
			return 1
		} else if a.(float32) == b.(float32){
			return 0
		} else {
			return -1
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) == b.(float64){
			return 0
		} else {
			return -1
		}
	}
	return 0
}