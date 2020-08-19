package segWordGetSlice

import (
	"fmt"
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg gse.Segmenter
	posSeg pos.Segmenter

	new = gse.New("zh,testdata/test_dict3.txt", "alpha")

	text = "你好世界, 山达尔星联邦百年剑道万古枯共和国联邦政府，天不生我李淳罡我的国家是中华人民共和国天不生我李淳罡，百年剑道万古枯三台中学"
)

func cut() {

	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)
}

func posAndTrim(cut []string) {
	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	posSeg.WithGse(seg)
	po := posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimPos(po)
	fmt.Println("trim pos: ", po)
}

func cutPos() {
	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))

	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)
}

func segCut(inputStr string) []string {
	// 加载默认字典
	//seg.LoadDict()
	// 载入词典
	seg.LoadDict("./dictionary.txt")

	// 分词文本
	//tb := []byte("山达尔星联邦百年剑道万古枯共和国联邦政府，Young manareyouok白衣兵圣安慕希钢铁侠特仑苏天不生我特伦李淳罡我的国家是中华人民共和国天不生我李淳罡，百年剑道万古枯三台中学")
	tb := []byte(inputStr)
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
	/*人家原来的：
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(string(tb), true))
	*/
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(string(tb))) //记得去掉注释
	ret := seg.Slice(string(tb))
	fmt.Println("The ret = ", ret) //记得去掉注释
	/* 人家原来的这是 ：
	segments := seg.Segment(tb)
	// 处理分词结果
	fmt.Println(gse.ToString(segments))

	segments1 := seg.Segment([]byte(text))
	fmt.Println(gse.ToString(segments1, true))
	*/
	return ret
}

func SegWordGetSlice(inputStr string) []string {
	ret := segCut(inputStr)
	return ret
}
