package main

import (
	"errors"
	"fmt"
)

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
		if tmp.FailPtr == nil {
			fmt.Printf("data: %v,isEnd: %v,length: %v,FailPrt.data: %v,%v", string(tmp.data), tmp.isEnd, tmp.length, "null", "-->")
		} else {
			fmt.Printf("data: %v,isEnd: %v,length: %v,FailPrt.data: %v,%v", string(tmp.data), tmp.isEnd, tmp.length, string(tmp.FailPtr.data), "-->")
		}
		fmt.Println()
		for _, v := range tmp.Child {
			if v != nil {
				arr = append(arr, v)
			}
		}
	}
	fmt.Println()
}

// 构建失败指针
func (p *AcNode) ConstructFailNext() (e error) {
	root := p
	// 模拟队列
	var queue []*AcNode
	queue = append(queue, p)
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		if q == nil {
			return errors.New("Nil Panic")
		}
		//为这个节点的所有子节点创建失效指针
		for i := 0; i < 26; i++ {
			//tmp是q的子节点
			tmp := q.Child[i]
			if tmp == nil {
				continue
			}
			// 如果为根节点则FailPtr为自己
			if q.data == '/' {
				tmp.FailPtr = root
			} else {
				//先将失效指针指向root
				tmp.FailPtr = root
				// 找到q的失效指针
				qFailPtr := q.FailPtr
				for qFailPtr != nil {
					qFailPtrChild := qFailPtr.Child[tmp.data-'a']
					// 如果找到字符
					if qFailPtrChild != nil {
						tmp.FailPtr = qFailPtrChild
						break
					}
					//没有找到就继续找
					qFailPtr = qFailPtr.FailPtr
				}
			}
			// 将子节点压入队列
			queue = append(queue, tmp)
		}
	}
	return
}

func (p *AcNode) MatchString(str string) error {
	root := p
	for i := 0; i < len(str); i++ {
		// 无法匹配就找到一个可以与之匹配的节点
		for p.Child[str[i]-'a'] == nil && p != root {
			p = p.FailPtr
		}

		p = p.Child[str[i]-'a']
		//从新匹配下一个
		if p == nil {
			p = root
		}

		//检测指针为结尾的路径是否是模式串& 看是否可以与失败指针所在的串匹配
		tmp := p
		for tmp != root {
			if tmp.isEnd {
				leftIndex := i - tmp.length + 1
				fmt.Println("匹配到的字符串", str[leftIndex:i+1])
			}
			tmp = tmp.FailPtr
		}
	}
	return nil
}
func main() {
	ac := InitAcNode()

	ac.AddString("bcd")
	ac.AddString("cd")
	ac.AddString("abcd")
	e := ac.ConstructFailNext()
	fmt.Println(e)
	fmt.Println(ac.FailPtr)
	ac.LevelRange()
	fmt.Println(ac.MatchString("abcd"))
}
