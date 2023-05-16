package main

import (
	"testing"
)

func Test_MidFInd(t *testing.T) {

	arr := []uint{1, 2, 4}
	index := finMid(arr, 3, 0, len(arr)-1)
	if index != 2 {
		t.Fatal("index = ", index)
	}
}
