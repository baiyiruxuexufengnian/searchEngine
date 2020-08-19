package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func makeInterfaceList(data interface{}) []interface{} {
	var res []interface{}
	switch data.(type) {
	case []int:
		for i := 0; i < len(data.([]int)); i++ {
			res = append(res, data.([]int)[i])
		}
	}
	return res
}

func TestBinarySearch(t *testing.T) {
	fmt.Println("------------------------------------------TestBinarySearch Start-----------------------------------------------------")
	type args struct {
		keys []interface{}
		key  interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "left edge", args: args{keys: makeInterfaceList([]int{1, 2, 3, 4, 5, 6, 7}), key: 0}, want: -1},
		{name: "right edge", args: args{keys: makeInterfaceList([]int{1, 2, 3, 4, 5, 6, 7}), key: 8}, want: -8},
		{name: "middle &not found", args: args{keys: makeInterfaceList([]int{1, 2, 3, 8, 9, 10, 11}), key: 4}, want: -4},
		{name: "find", args: args{keys: makeInterfaceList([]int{1, 2, 3, 8, 9, 10, 11}), key: 3}, want: 2},
	}

	BinarySearch := generateKeyBinarySearchFunc(nil, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.keys, tt.args.key, 7); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Println("------------------------------------------TestBinarySearch END-----------------------------------------------------")
}

//func Test_bPlusTree_Insert(t *testing.T) {
//	fmt.Println("------------------------------------------Test_bPlusTree_Insert Start-----------------------------------------------------")
//	order := 5
//	type Goods struct {
//		price int
//		name  string
//	}
//	//m := make(map[string]int)
//	tree := InitBPlusTree(order, nil, Goods{}.name)
//	//for i := 0; i < 3; i++ {    //生成随机的Key-->num
//	//	num := rand.Int() % 1000
//	//	for {
//	//		if _, ok := m[num]; ok {
//	//			num = rand.Int() % 1000
//	//		} else {
//	//			m[num] = true
//	//			break
//	//		}
//	//	}
//	//	// Value是一个Goods结构体类型的指针，内部有两个成员，成员类型是：int、string
//	//	//tree.Insert(num, &Goods{num, "test" + strconv.Itoa(num)})
//	//	insertVal := num
//	//	fmt.Println("insertVal = ", insertVal)
//	//	tree.Insert(num,insertVal)
//	//}
//	//target := "It was papa made me leave. I am a boy and I must obey him."
//	//m_map := SetFileId.SetFileId("HelloWord", "./test.txt,./Old_Man_And_Sea.txt,hello.txt,")
//	m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
//	for k, v := range m_map {
//		tree.Insert(k, v)
//	}
//
//	//m["Go"] = 1
//	//m["Python"] = 2
//	//m["奇安信"] = 3
//	//tree.Insert("Go", 1)
//	//tree.Insert("Python", 2)
//	//tree.Insert("奇安信", 3)
//
//	//fmt.Println("[ ", k, ":", v, " ]")
//	//fmt.Println("///////////////////////")
//	//price := tree.Search("HelloWord")
//	//fmt.Println("[ ", k, ":", v, " ]")
//	//fmt.Println("Search Result : ********", price, "***********")
//	//m := map[string][]int {"HelloWord":{1,2,3}}
//	fmt.Println("开始搜索关键词.......")
//	for k, _ := range m_map {
//		price := tree.Search(k)
//		//fmt.Println("[ ", k, ":", v, " ]")
//		fmt.Println("搜索词 ： ", k, "| 所在文件ID ： ", price)
//		//fmt.Println("Search Result : ********", price, "***********")
//		//if  price != v {
//		//	t.Errorf("Search() = %v, want %v", price, v)
//		//}
//	}
//	fmt.Println("------------------------------------------Test_bPlusTree_Insert END-----------------------------------------------------")
//}

func Test_bPlusTree_Delete(t *testing.T) {
	fmt.Println("------------------------------------------Test_bPlusTree_Delete Start-----------------------------------------------------")
	order := 5
	type Goods struct {
		price int
		name  string
	}
	m := make(map[int]bool)
	tree := InitBPlusTree(order, nil, Goods{}.price)
	for i := 0; i < 10; i++ {
		num := rand.Int() % 1000
		for {
			if _, ok := m[num]; ok {
				num = rand.Int() % 1000
			} else {
				m[num] = true
				break
			}
		}
		tree.Insert(num, &Goods{num, "test" + strconv.Itoa(num)})
	}
	tree.PrintSimply()
	for val, _ := range m {
		fmt.Println("delete: " + strconv.Itoa(val))
		tree.Delete(val)
		tree.PrintSimply()
		fmt.Println("alfter delete ")
	}
	fmt.Println("------------------------------------------Test_bPlusTree_Delete END-----------------------------------------------------")
}

func Test_bPlusTree_Insert(t *testing.T) {
	fmt.Println("------------------------------------------Test_bPlusTree_Insert Start-----------------------------------------------------")
	AddIndexToBtree()  //添加索引进入B+树
	fmt.Println("请输入搜索词 ： | 输入 ： END 结束")

	input := "桃花剑圣邓太阿，白衣兵圣陈芝豹"
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(stdin, &input)
	fmt.Printf("fmt.Fscan: %q\r\n", input)

	if input == "END" {
		fmt.Println("用户结束输入或退出")
	}
	result := keySearch(input)
	fmt.Printf("搜索词 ：%s,  | 所在文件ID ： %d\n", input, result)
	fmt.Println("------------------------------------------Test_bPlusTree_Insert END-----------------------------------------------------")
}
//
//func Test_input(t *testing.T) {
//	var input string
//	_, err := fmt.Scanf("%s\n", input)
//	if err != nil {
//		fmt.Println("hslasfdlfh")
//	}
//	fmt.Println(input)
//}
