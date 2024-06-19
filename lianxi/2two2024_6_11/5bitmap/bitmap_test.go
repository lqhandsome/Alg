package bitmap

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

var b *BitMap

func TestMain(m *testing.M) {
	b = NewBitMap(10)
	log.Info(b.Add(0))
	log.Info(b.Add(1))
	log.Info(b.Add(8))
	log.Info(b.Add(80))
	m.Run()
}

func Test_Add(t *testing.T) {

}

func Test_Get(t *testing.T) {
	log.Info(b.Get(0))
	log.Info(b.Get(1))
	log.Info(b.Get(8))
	log.Info(b.data)
}
