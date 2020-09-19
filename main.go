package main

import (
	"fmt"
)

func main() {

	var word string
	for {
		fmt.Println("请输入查找的单词：")
		fmt.Scanln(&word)
		fmt.Printf("查找结果：%t\n", tree.Search(word))
	}
}
