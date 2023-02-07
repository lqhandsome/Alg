package main

import "fmt"

func main() {
	manStr := "aaaaaba"
	//return
	fmt.Println(BMStringMatch(manStr, "ba"))
	fmt.Println(BMStringMatch(manStr, "a"))
	fmt.Println(BMStringMatch(manStr, "b"))
	fmt.Println(BMStringMatch(manStr, "ab"))
}

// 坏字符串和索引映射
func bm(str string) map[byte]int {
	ma := map[byte]int{}
	for k := range str {
		ma[str[k]] = k
	}
	return ma
}

// BMStringMatch
// BMStringMatch manStr 主串
// str 模式串
// 返回所在索引找不到返回-1
func BMStringMatch(manStr string, str string) int {
	bm := bm(str)
	for i := 0; i <= len(manStr)-len(str); {
		// 匹配
		var j int
		for j = len(str) - 1; j >= 0; j-- {
			if str[j] == manStr[j+i] {
				continue
			} else {
				break
			}
		}
		if j < 0 {
			return i
		}
		// 在模式串中可以匹配到坏字符串
		if v, ok := bm[manStr[j+i]]; ok {
			// 如果倒着滑动，强制往前滑
			if (j - v) > 0 {
				i = i + (j - v)
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return -1
}
