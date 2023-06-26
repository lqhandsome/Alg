package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestHeap(t *testing.T) {
	h := New(10)
	log.Info(h.add(10))
	log.Info(h.add(30))
	log.Info(h.add(44))
	log.Info(h.add(15))
	log.Info(h.add(9))
	log.Info(h.add(7))
	log.Info(h.add(7))
	log.Info(h.data)
}
