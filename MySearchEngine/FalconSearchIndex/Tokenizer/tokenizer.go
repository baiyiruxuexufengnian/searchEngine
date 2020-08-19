package Tokenizer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getNext(str2 string) []int {
	len2 := len(str2)
	next := make([]int, len2)
	next[0] = -1
	next[1] = 0
	i := 2
	cn := 0
	for i < len2 {
		if str2[i - 1] == str2[cn] {
			cn ++
			next[i] = cn
			i ++
		}else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i ++
		}
	}
	return next
}
//KMP算法，返回字串的起始小标，如果没有返回-1
func KMP(str1, str2 string) int {
	ret := -1
	len1 := len(str1)
	len2 := len(str2)
	if len1 == 0 || len2 == 0 || len1 < len2 {
		return -1
	}
	next := getNext(str2)
	var i1, i2 = 0, 0
	for i1 < len1 && i2 < len2 {
		if str1[i1] == str2[i2] {
			i1 ++
			i2 ++
		}else if next[i2] == -1 {
			i1 ++
		}else {
			i2 = next[i2]
		}
	}
	if i2 == len2 {
		ret = i1 - i2 + 1
	} else {
		ret = -1
	}
	return ret
}

func ReadBuffByIo(inputString string, fileName string) int {
	res := -1
	count := 0
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file, error message is : %v\n", err)
		return res
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("=========End of file reading=======")
			break
		}
		if err != nil {
			fmt.Printf("Failed to read file, error message is : %v\n", err)
			break
		}
		//fmt.Println(string(line))
		res :=  KMP(inputString, string(line))
		//fmt.Println("KMP res : ", res)
		if res > -1 {
			count ++
		}
		//fmt.Println(string(line))
	}
	return count
}