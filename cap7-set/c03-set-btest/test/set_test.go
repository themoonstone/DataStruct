package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"themoonstone/DataStruct/cap7-set/c03-set-btest/SetImplement"
	"unicode"
)

//  2         857209600 ns/op
// 2         871671950 ns/op
func BenchmarkLinkListSet(b *testing.B) {
	var list_set SetImplement.LinkedListSet
	list_set.Constructor()
	// 做一个单词统计应用
	file_name := "../../pride-and-prejudice.txt"
	// 读取文件内容
	f, err := os.Open(file_name)
	if nil != err {
		b.Errorf("open file failed! %v\n", err)
	}
	content, err := ioutil.ReadAll(f)
	if nil != err {
		b.Errorf("read file failed! %v\n", err)
	}
	// 提取单词
	words := strings.FieldsFunc(string(content), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	// 统计单词数量
	fmt.Printf("Total counts : %d\n", len(words))

	for i := 0; i < b.N; i++ {
		// 将所有单词存入集合中
		for _, word := range words {
			list_set.Add(word)
		}
		list_set.GetSize()
	}
}

// 30          46177460 ns/op
// 30          47173083 ns/op
func BenchmarkBSTreeSet(b *testing.B) {
	var bst_set SetImplement.BSTSet
	bst_set.Constructor()
	// 做一个单词统计应用
	file_name := "../../pride-and-prejudice.txt"
	// 读取文件内容
	f, err := os.Open(file_name)
	if nil != err {
		b.Errorf("open file failed! %v\n", err)
	}
	content, err := ioutil.ReadAll(f)
	if nil != err {
		b.Errorf("read file failed! %v\n", err)
	}
	// 提取单词
	words := strings.FieldsFunc(string(content), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	// 统计单词数量
	fmt.Printf("Total counts : %d\n", len(words))

	for i := 0; i < b.N; i++ {
		// 将所有单词存入集合中
		for _, word := range words {
			bst_set.Add(word)
		}
		bst_set.GetSize()
	}
}
