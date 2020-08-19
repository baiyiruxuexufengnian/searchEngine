package main

import (
	"MySearchEngine/FalconSearchIndex/InvertFile"
)

//fmt.Println("------------------------------------------Test_bPlusTree_Insert Start-----------------------------------------------------")
//order := 5
//type Goods struct {
//	price int
//	name  string
//}
////m := make(map[string]int)
//tree := InitBPlusTree(order, nil, Goods{}.name)
//m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
//for k, v := range m_map {
//tree.Insert(k, v)
//}
//fmt.Println("开始搜索关键词.......")
//for k, _ := range m_map {
//price := tree.Search(k)
//fmt.Println("搜索词 ： ", k, "| 所在文件ID ： ", price)
//}

//func Add() {
//	order := 5
//	type Goods struct {
//		price int
//		name  string
//	}
//	tree := InitBPlusTree(order, nil, Goods{}.name)
//	m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
//	for k, v := range m_map {
//		tree.Insert(k, v)
//	}
//	fmt.Println("开始搜索关键词.......")
//	for k, _ := range m_map {
//		price := tree.Search(k)
//		fmt.Println("搜索词 ： ", k, "| 所在文件ID ： ", price)
//	}
//}

func AddIndexToBtree () *bPlusTree {
	order := 5
	type Goods struct {
		price int
		name  string
	}
	tree := InitBPlusTree(order, nil, Goods{}.name)
	m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
	for k, v := range m_map {
		tree.Insert(k, v)
	}
	return tree
}

func keySearch(Key interface{}) interface{} {
	tree := AddIndexToBtree()
	ret := tree.Search(Key)

	return ret
}

