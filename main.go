package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	m := md5.New()
	m.Write([]byte("hello"))
	a := m.Sum(nil)
	fmt.Printf("%x\r\n", a)
	newMd := md5.New()
	newMd.Reset()
	md5String := fmt.Sprintf("%x", md5.Sum([]byte("今天我来讲哈希算法")))
	fmt.Println((md5String), time.Until(start))

}
