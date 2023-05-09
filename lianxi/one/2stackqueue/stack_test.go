package main

import (
	"fmt"
	"testing"
)

var (
	arrStack  *ArrStack
	listStack *ListStack
)

func TestMain(m *testing.M) {
	arrStack = InitArrStack(5)
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

}
