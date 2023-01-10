package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var strBuilder = strings.Builder{}
	for i := 0; i < 10; i++ {
		strBuilder.WriteString(strconv.Itoa(i))
	}
	fmt.Println(strBuilder.String())
	fmt.Println(strBuilder.Len())
	fmt.Println(strBuilder.Grow)
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
