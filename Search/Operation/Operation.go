package Operation

import (
	"MySearchEngine/FalconSearchIndex/InvertFile"
	"day13/Btree"
	"fmt"
	"unicode/utf8"
)

func AddIndexToBtree () *Btree.BPlusTree {
	order := 100
	type Goods struct {
		price int
		name  string
	}
	tree := Btree.InitBPlusTree(order, nil, Goods{}.name)

	//m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
	//m_map := InvertFile.AddIndexFile("./indexFile.txt", "./test7.txt,./test8.txt,./test9.txt,./test10.txt,")
	m_map := InvertFile.AddIndexFile("./indexFile1.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
	for k, v := range m_map {
		tree.Insert(k, v)
	}
	return tree
}

func KeySearch(Key interface{}) interface{} {
	tree := AddIndexToBtree()
	ret := tree.Search(Key)

	return ret
}

func MyPrint(input string) {
	fmt.Print("+------+")
	for i := 0; i < utf8.RuneCountInString(input) * 2; i ++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	fmt.Print("----------+")
	fmt.Println("---------+")
}


