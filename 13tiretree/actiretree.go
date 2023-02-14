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
	acNode := &AcNode{
		data: '/',
	}
	return acNode
}

// 插入一个模式串
func (p *AcNode) AddString(str string) (e error) {
	for i := 0; i < len(str); i++ {
		if p.Child[str[i]-'a'] == nil {
			tmp := &AcNode{
				data: str[i],
			}
			p.Child[str[i]-'a'] = tmp
		}
		p = p.Child[str[i]-'a']
	}
	p.isEnd = true
	p.length = len(str)
	return e
}

// 层序遍历TireTree
func (p *AcNode) LevelRange() {
	if p == nil {
		return
	}
	// 模拟队列
	var arr []*AcNode
	arr = append(arr, p)
	for len(arr) > 0 {
		tmp := arr[0]
		arr = arr[1:]
		fmt.Print(string(tmp.data), tmp.isEnd, tmp.length, "-->")
		for _, v := range tmp.Child {
			if v != nil {
				arr = append(arr, v)
			}
		}
	}
	fmt.Println()
}

// 构建失败指针
func (p *AcNode) FailNext() (e error) {
	root := p
	// 模拟队列
	var queue []*AcNode
	queue = append(queue, p)
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		for i := 0; i < 26; i++ {
			tmp := q.Child[i]
			if tmp == nil {
				continue
			}
			// 如果为根节点则FailPtr为自己
			if tmp == root {
				tmp.FailPtr = root
			} else {
				ttmp := tmp.FailPtr
				for ttmp != nil {
					ttmpC := ttmp.Child[tmp.data-'a']
					if ttmpC != nil {
						tmp.FailPtr = ttmpC
						break
					}
					ttmpC = ttmpC.FailPtr
				}
				if ttmp == nil {
					tmp.FailPtr = root
				}
			}
			queue = append(queue, tmp)
		}
	}
	return
}
func main() {
	ac := InitAcNode()
	ac.AddString("abc")
	ac.AddString("abd")
	ac.LevelRange()
}
