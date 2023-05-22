package main

import (
	"container/list"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type El struct {
	key   string
	value interface{}
}
type hashList struct {
	cap  uint8
	len  uint
	data []*list.List
}

// initHashList 初始化hash
func initHashList(cap uint8) *hashList {
	if cap == 0 {
		cap = 10
	}

	return &hashList{
		cap:  cap,
		data: make([]*list.List, cap, cap),
	}
}
func (h *hashList) Set(key string, value interface{}) error {
	if uint(h.cap) < h.len+1 {
		return fmt.Errorf("full")
	}
	index := h.hashFunc(key)
	head := h.data[index]

	if head == nil {
		headList := list.New()
		headList.PushBack(El{
			key,
			value,
		})
		h.data[index] = headList
	} else {
		l := h.data[index]
		// 存在覆盖
		for e := l.Front(); e != nil; e = e.Next() {
			st, ok := e.Value.(El)
			if !ok {
				panic("error")
			}

			// 覆盖
			if st.key == key {
				st.value = value
				e.Value = st

				return nil
			}
		}

		l.PushBack(El{
			key,
			value,
		})
	}
	h.len++
	return nil
}

func (h *hashList) Get(key string) (value interface{}, e error) {

	index := h.hashFunc(key)
	l := h.data[index]
	if l == nil {
		log.Warnf("key: %v,not found", key)

		return
	}
	for e := l.Front(); e != nil; e = e.Next() {
		st, ok := e.Value.(El)
		if !ok {
			panic(e.Value)
		}

		if st.key == key {
			return st.value, nil
		}
	}

	return nil, fmt.Errorf("key: %v,not found", key)
}

// hashFunc 简单哈希函数返回key的uin8%h.cap编码
func (h *hashList) hashFunc(key string) uint8 {
	return key[0] % h.cap
}
