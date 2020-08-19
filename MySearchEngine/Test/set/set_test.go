package set_test

import (
	"MySearchEngine/Test/set"
	"fmt"
	"testing"
)

func Test_set(t *testing.T) {
	m_set := set.New()
	for i := 1; i < 10; i ++ {
		m_set.Add(i)
	}

	fmt.Println("Result : ", m_set.Has(2))
	fmt.Println("Result : ", m_set.Has(6))
}

func Test_string(t *testing.T) {
	target := "boat which caught three good fish the first week"
	m_set := set.New()
	//slice := []string{}
	str := ""
	for i := 0; i < len(target); i ++ {
		if target[i] != ' ' {
			str += string(target[i])
		}else {
			//fmt.Println(str)
			m_set.Add(str)
			str = ""
		}
	}
	list := m_set.List()
	for _, v := range list {
		fmt.Println("Res : ", v)
	}
}
