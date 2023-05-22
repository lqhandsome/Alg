package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

var (
	hl *hashList
)

func TestMain(m *testing.M) {
	hl = initHashList(5)
	m.Run()
}

func Test_HashSet(t *testing.T) {
	hl.Set("a", "a")
	hl.Set("a", "aa")
	hl.Set("b", "b")
	hl.Set("bb", "bb")
	hl.Set("bbb", "b")
	hl.Set("bbbb", "b")
	hl.Set("bbbbb", "b")
	log.Info(hl.Set("bbbbb", "b"))

	log.Info(hl.Get("a"))
	log.Info(hl.Get("b"))
	log.Info(hl.Get("c"))
	//l := hl.data[7]
	//for e := l.Front(); e != nil; e = e.Next() {
	//	st, ok := e.Value.(El)
	//	if !ok {
	//		log.Error(e.Value)
	//	}
	//	fmt.Println(st)
	//}
}
