package main

import "reflect"

func IndexOfByReflect(arr interface{}, value interface{}) int {
	arrValue := reflect.ValueOf(arr)
	length := arrValue.Len()
	for i := 0; i < length; i++ {
		if arrValue.Index(i).Interface() == value {
			return i
		}
	}
	return -1
}

func IndexOfInterface(arr []interface{}, value interface{}) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}

func IndexOfInterfacePacking(value interface{}, arr ...interface{}) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}
