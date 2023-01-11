package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(fmt.Println(string([]byte{97, 107, 87})))
	hashList, e := getHashList(10)
	if e != nil {
		panic(e)
	}
	hashList.Set("a", 100)
	hashList.Set("k", 200)
	hashList.Set("W", 300)
	hashList.Set("1", 111)
	hashList.Set("2", 222)

	//key := 'a' % 10
	//fmt.Println(hashList.Delete("k"))
	fmt.Println(hashList.Delete("a"))
	fmt.Println(hashList.Get("W"))
	fmt.Println(hashList.Get("W"))
	fmt.Println(hashList.Get("1"))
	fmt.Println(hashList.Get("2"))
}

type Node struct {
	Key   string // Key
	Value int
	Next  *Node
}

type HashValue struct {
	*Node
	Key   string // Key
	Has   bool   // 本节点是否储存数据了
	Value int    // 储存的值
}
type HashList struct {
	Cap      uint        //容量
	Count    uint        // 当前数量
	Arr      []HashValue // 存储数据
	hashFunc func(string, uint) (uint, error)
}

func getHashList(cap uint) (*HashList, error) {
	if cap == 0 {
		return nil, errors.New("cap is empty")
	}
	hashFunc := Hash
	hasList := &HashList{
		Cap:      cap,
		Arr:      make([]HashValue, cap),
		hashFunc: hashFunc,
	}
	for k := range hasList.Arr {
		hasList.Arr[k].Node = &Node{}
	}
	return hasList, nil
}

/*
 cap 容量
*/
func Hash(key string, cap uint) (uint, error) {
	if key == "" {
		return 0, errors.New("key is empty")
	}
	hashKey := key[0] % uint8(cap)
	return uint(hashKey), nil
}

func (p *HashList) Set(key string, value int) error {
	hasKey, e := p.hashFunc(key, p.Cap)
	if e != nil {
		fmt.Println("set fail,error: ", e.Error())
		return e
	}

	// 如果节点没有值
	if !p.Arr[hasKey].Has {
		p.Arr[hasKey].Value = value
		p.Arr[hasKey].Has = true
		p.Arr[hasKey].Key = key
	} else {
		head := p.Arr[hasKey].Node
		for head.Next != nil {
			head = head.Next
		}
		tmpNode := &Node{Value: value, Key: key}
		head.Next = tmpNode
	}
	p.Count++
	return nil
}
func (p *HashList) Get(key string) (int, error) {
	hasKey, e := p.hashFunc(key, p.Cap)
	if e != nil {
		fmt.Println("set fail,error: ", e.Error())
		return 0, e
	}
	if !p.Arr[hasKey].Has {
		return 0, errors.New("not Found")
	}
	hashValue := p.Arr[hasKey]
	if hashValue.Key == key {
		return hashValue.Value, nil
	} else {
		head := hashValue.Node.Next
		for head != nil {
			if head.Key == key {
				return head.Value, nil
			}
			head = head.Next
		}
		return 0, errors.New("not Found")
	}
}

// 删除一个元素
func (p *HashList) Delete(key string) error {
	hasKey, e := p.hashFunc(key, p.Cap)
	if e != nil {
		fmt.Println("set fail,error: ", e.Error())
		return e
	}
	if !p.Arr[hasKey].Has {
		return errors.New("not Found")
	}

	//删除的是数组上的value
	if p.Arr[hasKey].Key == key {
		if p.Arr[hasKey].Node.Next == nil {
			p.Arr[hasKey].Has = false
			p.Arr[hasKey].Key = ""
			p.Arr[hasKey].Value = 0
		} else {
			p.Arr[hasKey].Key = p.Arr[hasKey].Node.Next.Key
			p.Arr[hasKey].Value = p.Arr[hasKey].Node.Next.Value
			p.Arr[hasKey].Node.Next = p.Arr[hasKey].Node.Next.Next
		}
		p.Count--
		return nil
	} else {
		preHead := p.Arr[hasKey].Node
		head := p.Arr[hasKey].Node.Next
		for head != nil {
			if head.Key == key {
				preHead.Next = head.Next
				p.Count--
				return nil
			}
			head = head.Next
			preHead = preHead.Next
		}
		return errors.New("not Found")
	}
}
