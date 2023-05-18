package main

import (
	list "container/list"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type hashList struct {
	cap  uint
	len  uint
	data []hashList
}

func initHastList() *hashList {

	return &hashList{}
}
func (h *hashList) Set(key string, value interface{}) error {

	return nil
}

func (h *hashList) Get(key string) (value interface{}, e error) {

	return
}

func main() {
	l := list.New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	log.Info(l.Len())
	log.Info(l.PushFront("d"))
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

func hashFunc(str string) {}
