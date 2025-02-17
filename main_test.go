package main

import (
	"testing"
	"time"
)

const ArrLength = 500

var _arr []interface{}
var _uintArr []uint

func init() {
	_arr = make([]interface{}, ArrLength)
	_uintArr = make([]uint, ArrLength)
	for i := 0; i < ArrLength-1; i++ {
		_uintArr[i] = uint(time.Now().Nanosecond()%10 + 2)
		_arr[i] = _uintArr[i]
	}
	_arr[ArrLength-1] = uint(1)
	_uintArr[ArrLength-1] = uint(1)
}

func BenchmarkIndexOfInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//IndexOfInterface(_arr, uint(1))
	}
}

func BenchmarkIndexOfInterfacePacking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//IndexOfInterfacePacking(uint(1), _arr...)
	}
}

func indexOfUint(arr []uint, value uint) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}

func BenchmarkIndexOfUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		indexOfUint(_uintArr, uint(1))
	}
}

func BenchmarkIndexOfByReflectInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//IndexOfByReflect(_arr, uint(1))
	}
}

func BenchmarkIndexOfByReflectUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//IndexOfByReflect(_uintArr, uint(1))
	}
}
