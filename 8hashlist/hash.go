package main

import (
	"errors"
	"fmt"
	"math"
)

// 一致性哈希算法

const NodeCap = 2 << 8

// Data 数据
type Data struct {
	Key   int
	Value int
}

// Machine 机器
type Machine struct {
	ip   string
	Data []Data
}

type ConsistentHash struct {
	MachineNum      uint                       // 机器数量
	MachineHashFunc func(string) (uint, error) // 机器哈希函数
	DataHashFunc    func(int) (uint, error)    // 数据哈希函数
	Node            []*Machine                 // 节点信息
}

func main() {

	m1 := Machine{
		ip:   "a",
		Data: nil,
	}
	m2 := Machine{
		ip:   "b",
		Data: nil,
	}
	cd, e := GetConsistentHash(m1, m2)
	if e != nil {
		fmt.Println(e)
	}
	e = cd.AddData(Data{
		Key:   1,
		Value: 11,
	})
	if e != nil {
		fmt.Println(e)
	}
	e = cd.AddData(Data{
		Key:   2,
		Value: 22,
	})
	cd.AddData(Data{
		Key:   392,
		Value: 392392,
	})
	cd.AddData(Data{
		Key:   388,
		Value: 388388,
	})
	e = cd.AddData(Data{
		Key:   400,
		Value: 4000,
	})
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(cd.GetData(2))
	fmt.Println(cd.GetData(1))
	fmt.Println(cd.GetData(3))
	fmt.Println(cd.GetData(392))
	fmt.Println(cd.GetData(388))
	fmt.Println(cd.GetData(400))
	fmt.Println(cd.DeleteData(2))
	fmt.Println(cd.GetData(2))
	fmt.Println(cd.GetData(3))
	fmt.Println(cd.GetData(1))
}
func HashOne(key string) (hashKey uint, e error) {
	if key == "" {
		fmt.Println("key is empty")
		return 0, errors.New("key is empty")
	}
	hashKey = uint((int(key[0]) << 2) % NodeCap)
	fmt.Printf("hashOne key: %v,hashKey: %v\r\n", key, hashKey)
	return
}

//hash 数据key
func HashTwo(key int) (hashKey uint, e error) {
	hash := key % NodeCap
	hash = int(math.Abs(float64(hash)))
	fmt.Printf("HashTwo key: %v,hashKey: %v\r\n", key, hash)
	return uint(hash), nil
}

func GetConsistentHash(machines ...Machine) (ch *ConsistentHash, e error) {
	if len(machines) == 0 {
		fmt.Println("machines is empty")
		return nil, errors.New("machines is empty")
	}
	node := make([]*Machine, NodeCap)
	for _, v := range machines {
		hashKey, e := HashOne(v.ip)
		if e != nil {
			return nil, e
		}
		tmp := v
		node[hashKey] = &tmp
	}
	ch = &ConsistentHash{
		MachineNum:      uint(len(machines)),
		Node:            node,
		MachineHashFunc: HashOne,
		DataHashFunc:    HashTwo,
	}
	return
}

// AddMachine 添加机器
func (p *ConsistentHash) AddMachine(machine Machine) error {
	hashKey, e := p.MachineHashFunc(machine.ip)
	if e != nil {
		return e
	}
	if p.Node[hashKey] != nil {
		fmt.Println("散列冲突", machine.ip, "已存在")
		return errors.New("散列冲突，机器已存在")
	}
	p.Node[hashKey] = &machine

	// 迁移数据
	//Step 1 查找下一个节点
	
	//Step 2 遍历下一个节点上的数据，发现需要迁移迁移到新机器上

	return nil
}

// DeleteMachine 删除机器
func (p *ConsistentHash) DeleteMachine(machine Machine) error {
	return nil
}

func (p *ConsistentHash) AddData(data Data) error {
	dataHashKey, e := p.DataHashFunc(data.Key)
	if e != nil {
		return e
	}
	for _, v := range p.Node[dataHashKey:] {
		if v != nil {
			v.Data = append(v.Data, data)
			return nil
		}
	}
	for _, v := range p.Node[0:dataHashKey] {
		if v != nil {
			v.Data = append(v.Data, data)
			return nil
		}
	}
	return errors.New("not found machine")
}
func (p *ConsistentHash) GetData(key int) (value int, e error) {
	dataHashKey, e := p.DataHashFunc(key)
	if e != nil {
		return 0, e
	}
	for _, v := range p.Node[dataHashKey:] {
		if v != nil {
			for _, vv := range v.Data {
				if vv.Key == key {
					return vv.Value, nil
				}
			}
			return 0, errors.New("not found")
		}
	}
	for _, v := range p.Node[:dataHashKey] {
		if v != nil {
			for _, vv := range v.Data {
				if vv.Key == key {
					return vv.Value, nil
				}
			}
			return 0, errors.New("not found")
		}

	}
	return 0, errors.New("not found")
}
func (p *ConsistentHash) DeleteData(key int) error {
	dataHashKey, e := p.DataHashFunc(key)
	if e != nil {
		return e
	}
	for _, v := range p.Node[dataHashKey:] {
		if v != nil {
			for kk, vv := range v.Data {
				if vv.Key == key {
					v.Data = append(v.Data[0:kk], v.Data[kk+1:]...)
					return nil
				}
			}
			return errors.New("not found")
		}
	}
	for _, v := range p.Node[:dataHashKey] {
		if v != nil {
			for kk, vv := range v.Data {
				if vv.Key == key {
					v.Data = append(v.Data[0:kk], v.Data[kk+1:]...)
				}
			}
			return errors.New("not found")
		}
	}
	return errors.New("not found")
}
