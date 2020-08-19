package SetFileId

import (
	"fmt"
	"testing"
)

func Test_AnalysisFile(t *testing.T)  {
	ret := AnalysisFile("./test.txt,./Old_Man_And_Sea.txt,")
	for _, v := range ret {
		fmt.Println("v : ", v)
	}
}

//func Test_fileID(t *testing.T) {
//	m_map := GetFileIdMap("boat which caught three good fish the first week", "./test.txt,./Old_Man_And_Sea.txt,")
//	for k,v := range m_map {
//		fmt.Println("k : ", k, "v : ", v)
//	}
//}

func Test_SetFileId(t *testing.T) {
	m_map := SetFileId("It was papa made me leave. I am a boy and I must obey him.", "./test.txt,./Old_Man_And_Sea.txt,")
	for k,v := range m_map {
		fmt.Println("k : ", k, "v : ", v)
	}
}

func Test_GetFileIdMap(t *testing.T) {
	m_map := GetFileIdMap("It was papa made me leave. I am a boy and I must obey him.", "./test.txt,./Old_Man_And_Sea.txt,")
	fmt.Println("Key_offset : ", m_map)
}

func Test_KeywordOccurrenceTimesInFile(t *testing.T)  {
	ret := KeywordOccurrenceTimesInFile("boat which caught three good fish the first week", "./test.txt")
	fmt.Println(ret)
}
