package InvertFile

import (
	"MySearchEngine/FalconSearchIndex/SetFileId"
	"MySearchEngine/FalconSearchIndex/Tokenizer"
	"bufio"
	"fmt"
	"io"
	"os"
)

func AddIndexFile(dictionary, files string) map[string][]int {
	ret := make(map[string][]int)
	fileSet := SetFileId.AnalysisFile(files)
	file, err := os.Open(dictionary)
	if err != nil {
		fmt.Printf("Failed to open file, error message is : %v\n", err)
		return ret
	}

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("=========End of dictionary file reading=======")
			break
		}
		if err != nil {
			fmt.Printf("Failed to read file, error message is : %v\n", err)
			break
		}
		slice := make([]int, 0)
		for i, v := range fileSet {
			//fmt.Println("i = ", i, "v = ", v)
			fileInSet, err := os.Open(v)

			if err != nil {
				fmt.Printf("The Failed to open file in fileSet, fileNum = %d error message is : %v\n", i, err)
				continue
			}
			ReaderInFileSet := bufio.NewReader(fileInSet)
			for {
				lineInFileSet, _, err := ReaderInFileSet.ReadLine()
				if err == io.EOF {
					//fmt.Println("=========End of FileSet reading=======")
					break
				}
				if err != nil {
					fmt.Printf("Failed to read file in FileSet, error message is : %v\n", err)
					break
				}
				occu_count := Tokenizer.KMP(string(lineInFileSet), string(line));
				if  occu_count > -1 {
					slice = append(slice, i)
					ret[string(line)] = slice
					break
				}
			}

			fileInSet.Close()
		}
		slice = make([]int, 0)
	}
		//fmt.Println(string(line))

	err = file.Close()

	if err != nil {
		fmt.Printf("Failed to close file error message is : %v\n", err)
	}
	return ret
}
