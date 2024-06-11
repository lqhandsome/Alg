package heap

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

var h *heap

func TestMain(m *testing.M) {
	h = NewHeap(10, []int{100, 22, 44, 5, 2, 1, 66})
	h.Print()
	log.Info("m.run")
	m.Run()
}

func Test_InitHeap(t *testing.T) {

	h.Print()
}

func Test_Remove(t *testing.T) {
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	log.Info(h.Remove())
	h.Print()
}

func BenchmarkIndexOfInterface(b *testing.B) {
	log.Info(time.Now())
}
