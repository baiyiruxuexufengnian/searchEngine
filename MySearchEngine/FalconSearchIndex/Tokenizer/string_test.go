package Tokenizer

import (
	"fmt"
	"testing"
)

func Test_kmp(t *testing.T) {
	ret := KMP("boat which caught three good fish the first week", "boat which caught three good fish the first week")
	fmt.Println(ret)
}

func Test_ReadBuffByIo(t *testing.T)  {
	ret := ReadBuffByIo("boat which caught three good fish the first week", "./test.txt")
	fmt.Println(ret)
}
