package main

import (
	"fmt"
)

type trie struct {
	data  uint8
	child [26]*trie
	cap   uint
}

// InitTire 初始化trie
func InitTire() *trie {
	root := &trie{
		data:  '/',
		child: [26]*trie{},
		cap:   26,
	}
	return root
}

// Insert 插入一个字符
func (p *trie) Insert(data string) error {
	tmp := p
	for i := 0; i < len(data); i++ {
		u := data[i] % 26

		// nil
		if tmp.child[u] == nil {
			tmp.child[u] = &trie{
				data:  data[i],
				child: [26]*trie{},
				cap:   0,
			}
		}
		tmp = tmp.child[u]
	}
	return nil
}

// GetString 查询字符串是否存在
func (p *trie) GetString(data string) bool {
	tmp := p
	for i := 0; i < len(data); i++ {
		u := data[i] % 26

		// nil
		if tmp.child[u] == nil {
			return false
		}
		tmp = tmp.child[u]
	}
	return true
}

// Range 打印tire数据
func (p *trie) Range() {
	if p == nil {
		return
	}
	tmp := p

	var queue []*trie
	queue = append(queue, tmp)
	queue = append(queue, nil)
	for len(queue) > 0 {
		tmpTrie := queue[0]
		queue = queue[1:]
		if tmpTrie == nil {
			fmt.Println()
			continue
		} else {
			fmt.Printf("%s -> ", string(tmpTrie.data))
		}
		for i := 0; i < 26; i++ {

			if tmpTrie.child[i] != nil {
				queue = append(queue, tmpTrie.child[i])
			}
			if i == 25 {
				queue = append(queue, nil)
			}
		}
		//fmt.Println()
	}
}
