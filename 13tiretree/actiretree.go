package main

import "fmt"

type AcNode struct {
	data    byte        // 当前节点数据
	Child   [26]*AcNode // 子节点
	isEnd   bool        // 是否结尾
	FailPtr *AcNode     // 失败指针
	length  int         // 当isEnd =true时,记录模式串长度
}

func InitAcNode() *AcNode {

	acNode := &AcNode{}

	fmt.Println(acNode.Child[10])
	return acNode
}

func main() {
	fmt.Println(InitAcNode())
}
