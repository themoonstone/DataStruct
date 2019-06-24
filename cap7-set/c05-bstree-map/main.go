package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"themoonstone/DataStruct/cap7-set/c05-bstree-map/map/mapImplement"
	"unicode"
)

func main() {
	var bm mapImplement.BasicTreeMap
	// map初始化
	bm.Constructor()
	// 做一个单词统计应用
	file_name := "../pride-and-prejudice.txt"
	// 读取文件内容
	f, err := os.Open(file_name)
	if nil != err {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if nil != err {
		panic(content)
	}
	 //提取单词
	words := strings.FieldsFunc(string(content), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	//词频统计
	for _, word := range words {
		if bm.Contains(word) {
			bm.Set(word, bm.Get(word).(int) + 1)
		} else {
			bm.Add(word,1)
		}
	}
	fmt.Println(bm.Get("what"))
	fmt.Println(bm.Size())
}
