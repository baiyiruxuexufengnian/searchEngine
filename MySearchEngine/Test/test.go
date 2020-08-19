package main

import (
	"MySearchEngine/Test/set"
	"fmt"
	"github.com/huichen/sego"
)

func targetIntoSetAndJudge(target []string, slice []string) {
	m_set := set.New()
	//slice := []string{}
	//str := ""
	//for i := 0; i < len(target); i ++ {
	//	if target[i] != ' ' {
	//		str += string(target[i])
	//	}else {
	//		//fmt.Println(str)
	//		m_set.Add(str)
	//		str = ""
	//	}
	//}
	for _, v := range target {
		m_set.Add(v)
	}
	for _, v := range(slice) {
		fmt.Println("The value = ", v, " | value:",v, "is in Set : ", m_set.Has(v))
	}
}

func stringToSlice(s string) []string {
	ss := ""
	slice := make([]string, 0)
	var index1 int = 0
	var index2 int = 0
	for index2 < len(s) {
		ss += string(s[index1])
		if s[index2] == ' '{
			slice = append(slice, ss)
			for index2 < len(s) && s[index2] == ' ' {
				index2 ++
			}
			ss = ""
		} else {
			index2 ++
		}
		index1 = index2
	}
	for _, v := range slice {
		fmt.Print("len=",len(v))
	}
	return slice
}

func segTostring(str string) string {
	s := ""
	//t := []string{}
	for i := 0; i < len(str); i ++ {
		//fmt.Println("************")
		if str[i] != '/' && str[i] != 'x' {
			s += string(str[i])
		}
	}
	return s
}


func main() {
	var segmenter sego.Segmenter
	//segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")
	segmenter.LoadDictionary("./Old_Man_And_Sea.txt")
	//target := []string{"boat", "which", "caught", "three", "good", "fish", "the", "first", "week"}
	target := "boat which caught three good fish the first week "
	text := []byte(target)
	segements := segmenter.Segment(text)
	str := sego.SegmentsToString(segements, false)
	//s := segTostring(str)
	//
	//slice1 := stringToSlice(target)
	//slice2 := stringToSlice(s)
	//
	//targetIntoSetAndJudge(slice1, slice2)
	//fmt.Println("Result : ", sego.SegmentsToString(segements, false))
	fmt.Println("Result : ", str)
}

//func main() {
//	var segmenter sego.Segmenter
//	segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")
//	//segmenter.LoadDictionary("./Old_Man_And_Sea.txt")
//	filename := "./Test/Text"
//	fp, err := os.Open(filename)
//	defer fp.Close()
//	if err != nil {
//		fmt.Println(filename, err)
//		return
//	}
//	buf := make([]byte, 4096)
//	n, _ := fp.Read(buf)
//	if n == 0 {
//		return
//	}
//
//	// 分词
//	segments := segmenter.Segment(buf)
//	// 处理分词结果
//	// 支持普通模式和搜索模式两种分词，见utils.go代码中SegmentsToString函数的注释。
//	// 如果需要词性标注，用SegmentsToString(segments, false)，更多参考utils.go文件
//	output := sego.SegmentsToSlice(segments, false)
//
//	//输出分词后的成语
//	for _, v := range output {
//		if len(v) == 12 {
//			fmt.Println(v)
//		}
//	}
//}





//package main
//
//import (
//	"fmt"
//	"github.com/huichen/sego"
//	"os"
//)
//
//func main() {
//
//	// 载入词典
//	var segmenter sego.Segmenter
//	segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")
//
//	//读取文件内容到buf中
//	filename := "./Test/Text"
//	fp, err := os.Open(filename)
//	defer fp.Close()
//	if err != nil {
//		fmt.Println(filename, err)
//		return
//	}
//	buf := make([]byte, 4096)
//	n, _ := fp.Read(buf)
//	if n == 0 {
//		return
//	}
//
//	// 分词
//	segments := segmenter.Segment(buf)
//	// 处理分词结果
//	// 支持普通模式和搜索模式两种分词，见utils.go代码中SegmentsToString函数的注释。
//	// 如果需要词性标注，用SegmentsToString(segments, false)，更多参考utils.go文件
//	output := sego.SegmentsToSlice(segments, false)
//
//	//输出分词后的成语
//	for _, v := range output {
//		if len(v) == 12 {
//			fmt.Println(v)
//		}
//	}
//}


