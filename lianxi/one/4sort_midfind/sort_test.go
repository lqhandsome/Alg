package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_popSort(t *testing.T) {

	arr := []uint{5, 1, 3, 5, 0}
	PopSort(arr)
	log.Info(arr)
}

func Test_SelectSort(t *testing.T) {

	arr := []uint{1}
	SelectSort(arr)
	log.Info(arr)
}

func Test_InsertSort(t *testing.T) {

	arr := []uint{1, 5, 3, 3, 0}
	InsertSort(arr)
	log.Info(arr)
}

func Test_QuickSort(t *testing.T) {

	arr := []uint{1}
	QuickSort(arr, 0, len(arr)-1)
	log.Info(arr)
}
