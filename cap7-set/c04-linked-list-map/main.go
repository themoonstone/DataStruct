package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"themoonstone/DataStruct/cap7-set/c04-linked-list-map/map/mapImplement"
	"unicode"
)

func main() {
	var lm mapImplement.LMap
	// map初始化
	lm.Constructor()
	// 做一个单词统计应用
	file_name := "../pride-and-prejudice.txt"
	// 读取文件内容
	f, err := os.Open(file_name)
	if nil != err {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if nil != err {
		panic(err)
	}
	// 提取单词
	words := strings.FieldsFunc(string(content), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	// 词频统计
	for _, word := range words {
		if lm.Contains(word) {
			lm.Set(word, lm.Get(word).(int)+1)
		} else {
			lm.Add(word, 1)
		}
	}

	fmt.Println(lm.Get("what"))
	fmt.Println(lm.Size())
}
