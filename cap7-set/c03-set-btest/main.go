package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"themoonstone/DataStruct/cap7-set/c03-set-btest/SetImplement"
	"unicode"
)

func main() {
	var list_set SetImplement.LinkedListSet
	list_set.Constructor()
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
	// 统计单词数量
	fmt.Printf("Total counts : %d\n", len(words))
	// 将所有单词存入集合中
	for _, word := range words {
		list_set.Add(word)
	}
	// 统计不重复的单词数量
	fmt.Printf("Different counts : %d\n", list_set.GetSize())
}
