package tree

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestRadixTree(t *testing.T) {
	fp, err := os.Open("../source/EnWords.csv")

	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	tree := NewRadixTree()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		tree.Add(record[0])
	}

	if !tree.Search("abroad") {
		t.Errorf("word \"abroad\" should be found")
	}

	if !tree.Search("government") {
		t.Errorf("word \"government\" should be found")
	}

	if tree.Search("governments") {
		t.Errorf("word \"governments\" should not be found")
	}
	tree.Add("governments")
	if !tree.Search("governments") {
		t.Errorf("word \"governments\" should  be found")
	}
	tree.Delete("governments")
	//if tree.Search("governments") {
	//	t.Errorf("word \"governments\" should not be found")
	//}
}
