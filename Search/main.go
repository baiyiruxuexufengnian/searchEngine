package main

import (
	"bufio"
	"day13/Operation"
	"day13/segWordGetSlice"
	"fmt"
	"os"
	"time"
)


func main() {
	fmt.Println("------------------------------------------添加索引以及搜索查询测试###开始###-----------------------------------------------------")
	startTime := time.Now()
	tree := Operation.AddIndexToBtree()  //添加索引进入B+树
	timeSpend := time.Since(startTime)
	fmt.Println("建立索引树所消耗的时间 : ", timeSpend)
	for {
		fmt.Println("请输入搜索词 ： | 输入 ： END 结束")

		input := "桃花剑圣邓太阿，白衣兵圣陈芝豹"
		stdin := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(stdin, &input)
		if err != nil {
			fmt.Println("标准输入有问题 ： ", err)
		}
		//fmt.Printf("fmt.Fscan: %q\r\n", input)
		if input == "END" {
			fmt.Println("用户结束输入或退出")
			break
		}
		strSlice := segWordGetSlice.SegWordGetSlice(input)
		startTime := time.Now()
		for _, v := range strSlice {
			//result := Operation.KeySearch(input)
			if len(v) >= 1 && len(v) <= 3 {
				continue
			}

			Operation.MyPrint(v)
			result := tree.Search(v)
			fmt.Printf("|搜索词|%s|所在文件ID|%d", v, result)
			fmt.Println("")
			Operation.MyPrint(v)
		}

		timeSpend := time.Since(startTime)
		fmt.Println("进行搜索操作所消耗的时间 : ", timeSpend)
	}
	fmt.Println("------------------------------------------添加索引以及搜索查询测试###结束###-----------------------------------------------------")
}
