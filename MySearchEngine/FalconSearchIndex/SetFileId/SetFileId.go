package SetFileId

import (
	"MySearchEngine/FalconSearchIndex/Tokenizer"
	"fmt"
)

func KeywordOccurrenceTimesInFile(inputString, file string) int {
	ret := Tokenizer.ReadBuffByIo(inputString, file)
	return ret
}
//./made.txt/,src/hello.txt
//拆分目录
func AnalysisFile(files string) []string {
	fileSlice := make([]string, 0)
	str := ""
	for i := 0; i < len(files); i ++ {
		if files[i] == ','{
			fileSlice = append(fileSlice, str)
			str = ""
		} else {
			str += string(files[i])
		}
	}
	return fileSlice
}

func SetFileId(inputString, files string) map[string][]int {
	ret := make(map[string][]int, 10)

	docid := 1
	slice := make([]int, 0)
	allFileName := AnalysisFile(files)

	for _, v := range allFileName {
		fmt.Println("fileSlice : ", v)
		Occu_count := Tokenizer.ReadBuffByIo(inputString, v)//KeywordOccurrenceTimesInFile(inputString, v)
		fmt.Println("Occu_count : ", Occu_count)
		if Occu_count > 0 {
			slice = append(slice, docid)
			fmt.Println("========38=============")
			ret[inputString] = slice
		}
		docid++
	}
	return ret
}

func GetFileIdMap(inputString, files string) map[string]int {
	t := SetFileId(inputString, files)

	offset := 0
	key_offset := make(map[string]int)
	for k, v := range t {
		fmt.Println("GetFileMap K : ", k, "v : ", v)
		offset = len(v)
		//fmt.Println(offset)
		key_offset[inputString] = offset
	}
	return key_offset
}
