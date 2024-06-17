package sort

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

var sli []int

func TestMain(m *testing.M) {
	sli = append(sli, 211, 3, 1, 56, 0, 21, 22, 22)
	m.Run()
}

func Test_popSort(t *testing.T) {
	popSort(sli)
	log.Info(sli)
}
func Test_selectSort(t *testing.T) {
	selectSort(sli)
	log.Info(sli)
}

func Test_insertSort(t *testing.T) {
	fmt.Println(len(`aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaqqqqqqqqqq`))
	insertSort(sli)
	log.Info(sli)
}
