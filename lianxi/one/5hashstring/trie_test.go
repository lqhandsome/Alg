package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

var tr *trie

func Test_InitTire(t *testing.T) {
	tr = InitTire()
}

func Test_Inert(t *testing.T) {
	tr = InitTire()
	e := tr.Insert("abde")
	e = tr.Insert("acde")
	if e != nil {
		log.Fatal(e)
	}
	log.Info(string(tr.child[19].data))

	tr.Range()

	log.Info(tr.GetString("abdec"))
}
