package main

import (
	"encoding/csv"
	"fmt"
	radixtree "github.com/mingjingc/go-data-structure/radixTree"
	"io"
	"os"
)

func main() {
	fp,err:=os.Open("EnWords.csv")

	if err!=nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	tree := radixtree.NewRadixTree()
	for {
		record,err:=reader.Read()
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Println(err)
			return
		}
		tree.Add(record[0])
	}
	var word string
	for {
		fmt.Println("请输入查找的单词：")
		fmt.Scanln(&word)
		fmt.Printf("查找结果：%t\n",tree.Search(word))
	}
}
