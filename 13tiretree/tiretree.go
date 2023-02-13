package main

import "fmt"

type TireTree struct {
	Data         byte
	Child        [26]*TireTree
	IsEndingChar bool
}

func InitTireTree() *TireTree {
	return &TireTree{
		Data:  '/',
		Child: [26]*TireTree{},
	}
}

// AddString
// 插入一个字符串
func (p *TireTree) AddString(str string) error {
	for i := 0; i < len(str); i++ {
		// 不存在则赋值
		if p.Child[str[i]-'a'] == nil {
			tmp := &TireTree{
				Data:  str[i],
				Child: [26]*TireTree{},
			}
			p.Child[str[i]-'a'] = tmp
		}
		//更换到子节点
		p = p.Child[str[i]-'a']
	}
	//判断是不是结尾字符
	p.IsEndingChar = true
	return nil
}

// 查找一个字符串
func (p *TireTree) Search(str string) bool {
	for i := 0; i < len(str); i++ {
		if p.Child[str[i]-'a'] != nil {
			p = p.Child[str[i]-'a']
		} else {
			return false
		}
	}
	if p.IsEndingChar {
		return true
	}
	return false
}

// 层序遍历TireTree
func (p *TireTree) LevelRange() {
	if p == nil {
		return
	}
	// 模拟队列
	var arr []*TireTree
	arr = append(arr, p)
	for len(arr) > 0 {
		tmp := arr[0]
		arr = arr[1:]
		fmt.Print(string(tmp.Data), "-->")
		for _, v := range tmp.Child {
			if v != nil {
				arr = append(arr, v)
			}
		}
	}
	fmt.Println()
}

//func main() {
//	tireTree := InitTireTree()
//	tireTree.LevelRange()
//	tireTree.AddString("abc")
//	tireTree.AddString("acd")
//	tireTree.AddString("ceg")
//	tireTree.LevelRange()
//
//	fmt.Println(tireTree.Search("abc"))
//	fmt.Println(tireTree.Search("c"))
//}
