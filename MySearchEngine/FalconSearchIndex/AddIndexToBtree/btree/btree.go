package main

import (
	"MySearchEngine/FalconSearchIndex/InvertFile"
	"fmt"
	"unsafe"
)

type NodeType int

const (
	LEAF_NODE NodeType = iota
	NON_LEAF_NODE
)

const (
	MAX_ORDER = 100
)

type bPlusTree struct {
	order int         //阶数
	root  interface{} // 可能是 treeLeafNode 或是 treeNonLeafNode
	//compareFunc  func(a, b interface{}) bool
	binarySearch func(keys []interface{}, key interface{}, size int) int
}

type nodeComm struct {
	parent      *treeNonLeafNode
	parentIndex int
	link        *link
	size        int //保存当前key的size
}

type treeLeafNode struct {
	nodeComm
	keys []interface{}
	data []interface{}
}

type treeNonLeafNode struct {
	nodeComm
	keys   []interface{}
	subPtr []interface{} //仅两种可能 1.treeLeafNode 2.treeNonLeafNode
}

var leafFieldOffset uintptr
var nonLeafFieldOffset uintptr

func getLeaf(l **link) *treeLeafNode {
	if l == nil {
		return nil
	}
	return (*treeLeafNode)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) - leafFieldOffset))
}

func getNonLeaf(l **link) *treeNonLeafNode {
	if l == nil {
		return nil
	}
	return (*treeNonLeafNode)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) - nonLeafFieldOffset))
}

func init() {
	dummy := &treeLeafNode{}
	leafFieldOffset = uintptr(unsafe.Pointer(&dummy.link)) - uintptr(unsafe.Pointer(dummy))
	nonLeafDummy := &treeNonLeafNode{}
	nonLeafFieldOffset = uintptr(unsafe.Pointer(&nonLeafDummy.link)) - uintptr(unsafe.Pointer(nonLeafDummy))
}

/*
 * 创建一颗指定阶数的B+树
 */
func InitBPlusTree(order int, compareFunc func(a, b interface{}) int, keyExample interface{}) *bPlusTree {
	if order < 3 || order > MAX_ORDER {
		panic("B+树的阶数不在范围内")
		return nil
	}
	binarySearch := generateKeyBinarySearchFunc(compareFunc, keyExample)
	return &bPlusTree{order, nil, binarySearch}
}

func newLeafNode(order int) *treeLeafNode {
	return &treeLeafNode{
		nodeComm{
			nil,
			-1,
			newLink(),
			0},
		make([]interface{}, order+1),
		make([]interface{}, order+1),}
	/*
	 * 为啥要make order+1个空间呢 ？
	 * 因为为了分裂方便
	 */
}

func newNonLeafNode(order int) *treeNonLeafNode {
	return &treeNonLeafNode{
		nodeComm{
			nil,
			-1,
			newLink(),
			0,
		},
		make([]interface{}, order),
		make([]interface{}, order+1),
	}
}


func (tree *bPlusTree) Insert(key, value interface{}) {
	if tree.root == nil {
		leaf := newLeafNode(tree.order)
		leaf.keys[0] = key
		leaf.data[0] = value
		leaf.size++
		tree.root = leaf
	} else {
		node := tree.root
		for {
			if leaf, ok := node.(*treeLeafNode); ok {
				// 如果是叶子节点
				tree.leafInsert(leaf, key, value)
				return
			} else {
				//非叶子节点
				index := tree.binarySearch(node.(*treeNonLeafNode).keys, key, node.(*treeNonLeafNode).size)
				if index >= 0 {
					node = node.(*treeNonLeafNode).subPtr[index]
				} else {
					node = node.(*treeNonLeafNode).subPtr[-index-1]
				}
			}
		}
	}
}

func (tree *bPlusTree) leafInsert(node *treeLeafNode, key, value interface{}) {
	insert := tree.binarySearch(node.keys, key, node.size)
	if insert > 0 {
		// TODO:insert大于0 说明有key重复 后期要处理
	} else {
		insert = -insert - 1
	}

	leafSimpleInsert(node, key, value, insert)
	if node.size > tree.order {
		split := node.size / 2
		// 创建
		rightLeaf := newLeafNode(tree.order)
		// 拷贝
		copyLeafNode(node, rightLeaf, split)
		// link
		addNext(&node.link, &rightLeaf.link)
		// 绑定父节点
		tree.leafBindParent(node, rightLeaf)
	}
}

func leafSimpleInsert(node *treeLeafNode, key, value interface{}, insert int) {
	for i := node.size; i > insert; i-- {
		node.keys[i] = node.keys[i-1]
		node.data[i] = node.data[i-1]
	}
	node.keys[insert] = key
	node.data[insert] = value
	node.size++
}

func copyLeafNode(ori, target *treeLeafNode, split int) {
	j := 0
	for i := split; i < len(ori.keys); i++ {
		target.keys[j] = ori.keys[i]
		target.data[j] = ori.data[i]
		j++
	}
	target.size = j
	ori.size = split
}

func (tree *bPlusTree) leafBindParent(left, right *treeLeafNode) {
	if left.parent == nil && right.parent == nil {
		parent := newNonLeafNode(tree.order)
		parent.keys[0] = right.keys[0]
		parent.subPtr[0] = left
		parent.subPtr[1] = right
		parent.size++

		left.parent = parent
		left.parentIndex = -1
		right.parent = parent
		right.parentIndex = 0

		tree.root = parent

	} else if left.parent != nil {
		insert := left.parentIndex + 1
		tree.nonLeafNodeInsert(left.parent, right.keys[0], right, insert)
	}
}

func (tree *bPlusTree) nonLeafNodeInsert(parent *treeNonLeafNode, key, treeNode interface{}, insert int) {
	for i := parent.size; i >= insert; i-- {
		if i == insert {
			parent.keys[i] = key
			parent.subPtr[i+1] = treeNode
			parent.size++
			if leaf, ok := treeNode.(*treeLeafNode); ok {
				leaf.parent = parent
				leaf.parentIndex = insert
			} else {
				treeNode.(*treeNonLeafNode).parent = parent
				treeNode.(*treeNonLeafNode).parentIndex = insert
			}
			break

		} else {
			parent.keys[i] = parent.keys[i-1]
			parent.subPtr[i+1] = parent.subPtr[i]
		}
		if leaf, ok := parent.subPtr[i+1].(*treeLeafNode); ok {
			leaf.parentIndex++
		} else {
			if parent.subPtr[i+1] == nil {
				continue
			}
			parent.subPtr[i+1].(*treeNonLeafNode).parentIndex++
		}
	}
	if parent.size == tree.order {
		split := tree.order / 2
		right := newNonLeafNode(tree.order)
		copyNonLeafNode(parent, right, split)
		addNext(&parent.link,&right.link)
		tree.nonLeafNodeBindParent(parent, right)
	}
}

func copyNonLeafNode(left, right *treeNonLeafNode, split int) {
	left.size = split
	j := 0
	for i := split; i < len(left.keys); i++ {
		right.keys[j] = left.keys[i]
		right.subPtr[j+1] = left.subPtr[i+1]
		if leaf, ok := right.subPtr[j+1].(*treeLeafNode); ok {
			leaf.parent = right
			leaf.parentIndex = j
		} else {
			right.subPtr[j+1].(*treeNonLeafNode).parent = right
			right.subPtr[j+1].(*treeNonLeafNode).parentIndex = j
		}
		j++
	}
	right.size = j
}

func (tree *bPlusTree) nonLeafNodeBindParent(left, right *treeNonLeafNode) {
	if left.parent == nil && right.parent == nil {
		parent := newNonLeafNode(tree.order)
		parent.keys[0] = right.keys[0]
		parent.subPtr[0] = left
		parent.subPtr[1] = right
		parent.size++

		right.parent = parent
		left.parent = parent
		right.parentIndex = 0
		left.parentIndex = -1

		tree.root = parent

	} else if left.parent != nil {
		insert := left.parentIndex + 1
		tree.nonLeafNodeInsert(left.parent, right.keys[0], right, insert)
	}
}


func (tree *bPlusTree) Delete(key interface{}) {
	node := tree.root
	if node == nil {
		return
	}
	for {
		if leaf, ok := node.(*treeLeafNode); ok {
			tree.deleteFormLeaf(key, leaf)
			return
		} else {
			nonLeaf, _ := node.(*treeNonLeafNode)
			index := tree.binarySearch(nonLeaf.keys, key, nonLeaf.size)
			if index >= 0 {
				node = nonLeaf.subPtr[index+1]
			} else {
				node = nonLeaf.subPtr[-index-1]
			}
		}
	}
}

func (tree *bPlusTree) deleteFormLeaf(key interface{}, leaf *treeLeafNode) {
	index := tree.binarySearch(leaf.keys, key, leaf.size)
	if index < 0 {
		return
	}
	simpleDelete(leaf, index)
	if leaf.parent == nil {
		if leaf.size == 0 {
			tree.root = nil
		}
		return
	}
	if leaf.size < (tree.order+1)/2 {
		left := getLeaf(leaf.link.pre)
		right := getLeaf(leaf.link.next)
		if choiceLeft := makeSiblingChoice(left, right); choiceLeft {
			//选择左边的
			if left.size+leaf.size <= tree.order {
				//可以合并
				tree.mergeToLeftLeaf(left, leaf)
			} else {
				//从左边借一个
				shiftFromLeftLeaf(left, leaf)
			}
		} else {
			//选择右边
			if right.size+leaf.size <= tree.order {
				//合并
				tree.mergeToRightLeaf(leaf, right)
			} else {
				//从右边借一个
				shiftFromRightLeaf(leaf, right)
			}
		}
	}

}

func simpleDelete(leaf *treeLeafNode, index int) {
	for i := index + 1; i < leaf.size; i++ {
		leaf.keys[i-1] = leaf.keys[i]
		leaf.data[i-1] = leaf.data[i-1]
	}
	leaf.size--
	if index == 0 && leaf.parent != nil && leaf.parentIndex != -1 {
		replaceRecursive(leaf.parent, leaf.parentIndex, leaf.keys[0])
	}
}

func makeSiblingChoice(left, right interface{}) bool {
	switch left.(type) {
	case *treeLeafNode:
		if left.(*treeLeafNode) == nil {
			return false
		}
		if right.(*treeLeafNode) == nil {
			return true
		}
		if left.(*treeLeafNode).size > right.(*treeLeafNode).size {
			return true
		}
		return false
	case *treeNonLeafNode:
		if left.(*treeNonLeafNode) == nil {
			return false
		}
		if right.(*treeNonLeafNode) == nil {
			return true
		}
		if left.(*treeNonLeafNode).size > right.(*treeNonLeafNode).size {
			return true
		}
		return false
	default:
		return false
	}
}

func shiftFromLeftLeaf(left, leaf *treeLeafNode) {
	for i := leaf.size; i > 0; i-- {
		leaf.keys[i] = leaf.keys[i-1]
		leaf.data[i] = leaf.data[i-1]
	}
	leaf.keys[0] = left.keys[left.size-1]
	leaf.data[0] = left.data[left.size-1]
	leaf.size++
	left.size--
	replaceRecursive(leaf.parent, leaf.parentIndex, leaf.keys[0])
}

func shiftFromRightLeaf(leaf, right *treeLeafNode) {
	leaf.keys[leaf.size] = right.keys[0]
	leaf.data[leaf.size] = right.data[0]
	leaf.size++
	right.size--
	for i := 0; i < right.size; i++ {
		right.keys[i] = right.keys[i+1]
		right.data[i] = right.data[i+1]
	}
	replaceRecursive(right.parent, right.parentIndex, right.keys[0])
}
func replaceRecursive(nonLeaf *treeNonLeafNode, index int, key interface{}) {
	nonLeaf.keys[index] = key
	if index == 0 && nonLeaf.parent != nil && nonLeaf.parentIndex != -1 {
		replaceRecursive(nonLeaf.parent, nonLeaf.parentIndex, key)
	}
}
func (tree *bPlusTree) mergeToLeftLeaf(left, leaf *treeLeafNode) {
	for i := 0; i < leaf.size; i++ {
		left.keys[left.size] = leaf.keys[i]
		left.data[left.size] = leaf.data[i]
		left.size++
	}
	leaf.link.deleteSelf()
	tree.deleteFromNonLeaf(leaf.parent, leaf.parentIndex)
}
func (tree *bPlusTree) mergeToRightLeaf(leaf, right *treeLeafNode) {
	for i := 0; i < right.size; i++ {
		leaf.keys[leaf.size] = right.keys[i]
		leaf.data[leaf.size] = right.data[i]
		leaf.size++
	}
	right.link.deleteSelf()
	tree.deleteFromNonLeaf(right.parent, right.parentIndex)
}

func (tree *bPlusTree) deleteFromNonLeaf(nonLeaf *treeNonLeafNode, delete int) {
	simpleDeleteFromNonLeaf(nonLeaf, delete)
	if nonLeaf.parent == nil {
		if nonLeaf.size == 0 {
			tree.root = nonLeaf.subPtr[0]
			switch nonLeaf.subPtr[0].(type) {
			case *treeLeafNode:
				tree.root.(*treeLeafNode).parent = nil
			case *treeNonLeafNode:
				tree.root.(*treeNonLeafNode).parent = nil
			}
		}
		return
	}
	if nonLeaf.size < (tree.order-1)/2 {
		left := getNonLeaf(nonLeaf.link.pre)
		right := getNonLeaf(nonLeaf.link.next)
		if choiceLeft := makeSiblingChoice(left, right); choiceLeft {
			if left.size+nonLeaf.size <= tree.order-1 {
				tree.mergeLeftNonLeaf(left, nonLeaf)
			} else {
				shiftFromLeftNonLeaf(left, nonLeaf)
			}
		} else {
			if right.size+nonLeaf.size <= tree.order-1 {
				tree.mergeRightNonLeaf(nonLeaf, right)
			} else {
				shiftFromRightNonLeaf(nonLeaf, right)
			}
		}
	}
}

func (tree *bPlusTree) mergeLeftNonLeaf(left, nonLeaf *treeNonLeafNode) {
	for i := 0; i < nonLeaf.size; i++ {
		left.keys[left.size] = nonLeaf.keys[i]
		left.subPtr[left.size+1] = nonLeaf.subPtr[i+1]
		switch left.subPtr[left.size+1].(type) {
		case *treeNonLeafNode:
			left.subPtr[left.size+1].(*treeNonLeafNode).parent = left
			left.subPtr[left.size+1].(*treeNonLeafNode).parentIndex = left.size
		case *treeLeafNode:
			left.subPtr[left.size+1].(*treeLeafNode).parent = left
			left.subPtr[left.size+1].(*treeLeafNode).parentIndex = left.size
		}
		left.size++
	}
	nonLeaf.link.deleteSelf()
	tree.deleteFromNonLeaf(nonLeaf.parent, nonLeaf.parentIndex)
}

func (tree *bPlusTree) mergeRightNonLeaf(nonLeaf, right *treeNonLeafNode) {
	for i := 0; i < right.size; i++ {
		nonLeaf.keys[nonLeaf.size] = right.keys[i]
		nonLeaf.subPtr[nonLeaf.size+1] = right.subPtr[i+1]
		switch nonLeaf.subPtr[nonLeaf.size+1].(type) {
		case *treeNonLeafNode:
			nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parent = nonLeaf
			nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parentIndex = nonLeaf.size
		case *treeLeafNode:
			nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parent = nonLeaf
			nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parentIndex = nonLeaf.size
		}
		nonLeaf.size++
	}
	right.link.deleteSelf()
	tree.deleteFromNonLeaf(right.parent, right.parentIndex)
}

func shiftFromRightNonLeaf(nonLeaf, right *treeNonLeafNode) {
	nonLeaf.keys[nonLeaf.size] = right.keys[0]
	nonLeaf.subPtr[nonLeaf.size+1] = right.subPtr[1]
	switch nonLeaf.subPtr[nonLeaf.size+1].(type) {
	case *treeNonLeafNode:
		nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parent = nonLeaf
		nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parentIndex = nonLeaf.size
	case *treeLeafNode:
		nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parent = nonLeaf
		nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parentIndex = nonLeaf.size
	}
	nonLeaf.size++
	right.size--
	for i := 0; i < right.size; i++ {
		right.keys[i] = right.keys[i+1]
		right.subPtr[i+1] = right.subPtr[i+2]
		switch right.subPtr[i+1].(type) {
		case *treeLeafNode:
			right.subPtr[i+1].(*treeLeafNode).parentIndex--
		case *treeNonLeafNode:
			right.subPtr[i+1].(*treeNonLeafNode).parentIndex--
		}
	}
	replaceRecursive(right.parent, right.parentIndex, right.keys[0])
}

func shiftFromLeftNonLeaf(left, nonLeaf *treeNonLeafNode) {
	for i := nonLeaf.size; i > 0; i-- {
		nonLeaf.keys[i] = nonLeaf.keys[i-1]
		nonLeaf.subPtr[i+1] = nonLeaf.subPtr[i]
		switch nonLeaf.subPtr[i].(type) {
		case *treeNonLeafNode:
			nonLeaf.subPtr[i].(*treeNonLeafNode).parentIndex++
		case *treeLeafNode:
			nonLeaf.subPtr[i+1].(*treeLeafNode).parentIndex++
		}
	}
	nonLeaf.size++
	left.size--
	nonLeaf.keys[0] = left.keys[left.size]
	nonLeaf.subPtr[1] = left.subPtr[left.size+1]
	switch nonLeaf.subPtr[1].(type) {
	case *treeLeafNode:
		nonLeaf.subPtr[1].(*treeLeafNode).parentIndex = 0
		nonLeaf.subPtr[1].(*treeLeafNode).parent = nonLeaf
	case *treeNonLeafNode:
		nonLeaf.subPtr[1].(*treeNonLeafNode).parentIndex = 0
		nonLeaf.subPtr[1].(*treeNonLeafNode).parent = nonLeaf
	}
	replaceRecursive(nonLeaf.parent, nonLeaf.parentIndex, nonLeaf.keys[0])
}

func simpleDeleteFromNonLeaf(nonLeaf *treeNonLeafNode, delete int) {
	nonLeaf.size--
	for i := delete; i < nonLeaf.size; i++ {
		nonLeaf.keys[i] = nonLeaf.keys[i+1]
		nonLeaf.subPtr[i+1] = nonLeaf.subPtr[i+2]
		switch nonLeaf.subPtr[i+1].(type) {
		case *treeLeafNode:
			nonLeaf.subPtr[i+1].(*treeLeafNode).parentIndex = i
		case *treeNonLeafNode:
			nonLeaf.subPtr[i+1].(*treeNonLeafNode).parentIndex = i
		}
	}
	if delete == 0 && nonLeaf.parent != nil && nonLeaf.parentIndex != -1 {
		replaceRecursive(nonLeaf.parent, nonLeaf.parentIndex, nonLeaf.keys[0])
	}
}



/*
 * 用于链接B➕树的同级节点
 */
type link struct {
	pre  **link
	next **link
}

/*
 * 初始化link, 感觉不需要循环链表
 */
func newLink() *link {
	l := &link{}
	l.pre = nil
	l.next = nil
	return l
}

/*
 * 把 add 添加到 l 后边
 */
func addNext(ori, add **link) {
	if (*ori).next == nil {
		(*ori).next = add
		(*add).pre = ori
	} else {
		(*(*ori).next).pre = add
		(*add).next = (*ori).next
		(*add).pre = ori
		(*ori).next = add
	}
}

func (l *link) deleteSelf() {
	if l.pre != nil {
		(*l.pre).next = l.next
	}
	if l.next != nil {
		(*l.next).pre = l.pre
	}
}


func (tree *bPlusTree) Search(key interface{}) interface{} {
	node := tree.root
	for {
		if leaf, ok := node.(*treeLeafNode); ok {
			index := tree.binarySearch(leaf.keys, key, leaf.size)
			if index >= 0 {
				return leaf.data[index]
			} else {
				return nil
			}
		} else {
			index := tree.binarySearch(node.(*treeNonLeafNode).keys, key, node.(*treeNonLeafNode).size)
			if index >= 0 {
				node = node.(*treeNonLeafNode).subPtr[index+1]
			} else {
				node = node.(*treeNonLeafNode).subPtr[-index-1]
			}
		}
	}
}



func (tree *bPlusTree) PrintSimply() {
	queue := make([]interface{}, 0)
	queue = append(queue, tree.root)
	child := make([]interface{}, 0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if leaf, ok := node.(*treeLeafNode); ok {
			for i := 0; i < leaf.size; i++ {
				fmt.Print(leaf.keys[i], " ")
			}
			fmt.Printf("=")
		} else if nonLeaf, ok := node.(*treeNonLeafNode); ok {

			for i := 0; i < nonLeaf.size; i++ {
				fmt.Print(nonLeaf.keys[i], " ")
			}
			fmt.Printf("=")
			for i := 0; i <= nonLeaf.size; i++ {
				child = append(child, nonLeaf.subPtr[i])
			}
		} else {
			continue
		}

		if len(queue) == 0 {
			fmt.Printf("\n")
			if len(child) == 0 {
				return
			}

			queue = child
			child = make([]interface{}, 0)
		}
	}
}





func generateKeyBinarySearchFunc(compareFunc func(a, b interface{}) int, keyExample interface{}) func(data []interface{}, key interface{}, size int) int {
	if compareFunc == nil {
		switch keyExample.(type) {
		case int:
			compareFunc = func(a, b interface{}) int {
				if a.(int) < b.(int) {
					return -1
				} else if a.(int) > b.(int) {
					return 1
				}
				return 0
			}
		case float64:
			compareFunc = func(a, b interface{}) int {
				if a.(float64) < b.(float64) {
					return -1
				} else if a.(float64) > b.(float64) {
					return 1
				}
				return 0
			}
		case string:
			compareFunc = func(a, b interface{}) int {
				if a.(string) < b.(string) {
					return -1
				} else if a.(string) > b.(string) {
					return 1
				}
				return 0
			}
		default:
			panic("请定义key比较规则")
		}
	}
	return func(keys []interface{}, key interface{}, size int) int {
		low := 0
		high := size - 1
		mid := low + (high-low)/2
		var res int
		for low <= high {
			res = compareFunc(keys[mid], key)
			if res == 0 {
				//找到了
				return mid
			} else if res > 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
			mid = low + (high-low)/2
		}
		//没找到 但是返回 最接近key且大于key的 位置下标 的相反数
		return -low - 1
	}
}




func Add() {
	order := 5
	type Goods struct {
		price int
		name  string
	}
	tree := InitBPlusTree(order, nil, Goods{}.name)
	m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
	for k, v := range m_map {
		tree.Insert(k, v)
	}
	fmt.Println("开始搜索关键词.......")
	for k, _ := range m_map {
		price := tree.Search(k)
		fmt.Println("搜索词 ： ", k, "| 所在文件ID ： ", price)
	}
}

func AddIndexToBtree () *bPlusTree {
	order := 5
	type Goods struct {
		price int
		name  string
	}
	tree := InitBPlusTree(order, nil, Goods{}.name)
	m_map := InvertFile.AddIndexFile("./dictionary.txt", "./test.txt,./test1.txt,./test2.txt,./test3.txt,./test4.txt,./test5.txt,")
	for k, v := range m_map {
		tree.Insert(k, v)
	}
	return tree
}

func keySearch(Key interface{}) interface{} {
	tree := AddIndexToBtree()
	ret := tree.Search(Key)

	return ret
}


