package main

import (
	"fmt"
	"math"
)

// 浏览器提词功能
// 计算两个字符串的编辑距离
var (
	distance int = math.MaxInt
	ijDist   [][]int
	st1      string
	st2      string
)

// 回溯方法计算莱文斯坦距离 lwstBT(st2, 0, len(st2), st1, 0, len(st1), 0)
func lwstBT(str1 string, i int, n int, str2 string, j, m int, dist int) {
	// 某个字符串遍历完成
	if i == n || j == m {
		if m > j {
			dist += m - j
		} else {
			dist += n - i
		}
		if dist < distance {
			distance = dist
		}
		return
	}
	if ijDist[i][n] == dist {
		return
	} else {
		ijDist[i][n] = dist
	}
	// 匹配
	if str1[i] == str2[j] {
		lwstBT(str1, i+1, n, str2, j+1, m, dist)
	} else { // 不匹配
		lwstBT(str1, i+1, n, str2, j, m, dist+1)   // 删除str1[i]或者str2[j]前添加一个字符
		lwstBT(str1, i, n, str2, j+1, m, dist+1)   // 删除str2[j]或者str1[i]前添加一个字符
		lwstBT(str1, i+1, n, str2, j+1, m, dist+1) // 将a[i]和b[j]替换为相同字符
	}
}

func lwstBTNoFilter(str1 string, i int, n int, str2 string, j, m int, dist int) {
	// 某个字符串遍历完成
	if i == n || j == m {
		if m > j {
			dist += m - j
		} else {
			dist += n - i
		}
		if dist < distance {
			distance = dist
		}
		return
	}
	// 匹配
	if str1[i] == str2[j] {
		lwstBT(str1, i+1, n, str2, j+1, m, dist)
	} else { // 不匹配
		lwstBT(str1, i+1, n, str2, j, m, dist+1)   // 删除str1[i]或者str2[j]前添加一个字符
		lwstBT(str1, i, n, str2, j+1, m, dist+1)   // 删除str2[j]或者str1[i]前添加一个字符
		lwstBT(str1, i+1, n, str2, j+1, m, dist+1) // 将a[i]和b[j]替换为相同字符
	}
}
func init() {
	fmt.Println(11)
}
func main() {

	lwstBT(st2, 0, len(st2), st1, 0, len(st1), 0)
	fmt.Println(distance)
}
