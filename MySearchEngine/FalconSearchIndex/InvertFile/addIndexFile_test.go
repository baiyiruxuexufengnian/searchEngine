package InvertFile

import (
	"fmt"
	"testing"
)

func Test_addIndexFile(t *testing.T) {
	/*./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,*/
	m_map := AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,v")
	for k, v := range m_map {
		fmt.Println("--------+-----------------------------+---------------")
		fmt.Println("搜索词 | ", k, " | 所在文件ID | ", v, " |")
	}
}
