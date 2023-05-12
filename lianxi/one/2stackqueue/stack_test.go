package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

var (
	arrStack  *ArrStack
	listStack *ListStack
)

func TestMain(m *testing.M) {
	arrStack = InitArrStack(5)
	listStack = InitListStack(5)
	m.Run()
}

func Test_ArrStack_Add(t *testing.T) {

	e := arrStack.Add("a")
	if e != nil {
		panic(e)
	}
	arrStack.Add("b")
	arrStack.Add("c")
	arrStack.Add("d")
	arrStack.Add("e")
	arrStack.Add("f")
	fmt.Printf("%+v\r\n", arrStack)
	arrStack.Range()
}

func Test_ArrStack_Pop(t *testing.T) {
	arrStack.Add("b")
	arrStack.Add("c")
	arrStack.Add("d")
	arrStack.Add("e")
	arrStack.Add("f")

	var data string
	var e error
	for !arrStack.IsEmpty() {
		data, e = arrStack.Pop()
		if e != nil {
			panic(e)
		}
		fmt.Println(data)
	}

}
func Test_ArrStack_Range(t *testing.T) {
	listStack.Range()
}

// Test_ListStack_Init 测试初始化ListStack
func Test_ListStack_Init(t *testing.T) {

	log.Info(listStack.head, listStack.tail)
}

func Test_ListStack_IsEmpty_IsFull(t *testing.T) {
	log.Info(listStack.IsEmpty(), listStack.IsFull())
}

func Test_ListStack_Add(t *testing.T) {
	e := listStack.Add("a")
	if e != nil {
		log.Fatal(e)
	}

	e = listStack.Add("b")
	if e != nil {
		log.Fatal(e)
	}

	e = listStack.Add("c")
	if e != nil {
		log.Fatal(e)
	}

	e = listStack.Add("d")
	if e != nil {
		log.Fatal(e)
	}
	e = listStack.Add("e")
	if e != nil {
		log.Fatal(e)
	}
	//e = listStack.Add("f")
	//if e != nil {
	//	log.Fatal(e)
	//}
	listStack.Range()
	log.Info(listStack.IsEmpty(), listStack.IsFull())
}

func Test_ListStack_Pop(t *testing.T) {
	e := listStack.Add("a")
	if e != nil {
		log.Fatal(e)
	}
	log.Info(listStack.Pop())
	log.Info(listStack.Pop())
}
