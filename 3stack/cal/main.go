package main

import (
	"fmt"
	"strconv"
)

func main() {

	strArr := []string{"2", "+", "3", "*", "5", "/", "5"}
	numArr := []int64{}
	chartArr := []string{}

	for _, v := range strArr {
		fmt.Println(numArr, chartArr, v)
		if v == "+" || v == "-" || v == "*" || v == "/" {
			if v == "+" || v == "-" {
				for len(chartArr) > 0 {
					var tmp int64
					if chartArr[len(chartArr)-1] == "-" {
						tmp = numArr[len(numArr)-2] - numArr[len(numArr)-1]
					} else {
						tmp = numArr[len(numArr)-2] + numArr[len(numArr)-1]
					}
					numArr = numArr[0 : len(numArr)-2]
					numArr = append(numArr, tmp)
				}
				chartArr = append(chartArr, v)
			} else {
				if chartArr[len(chartArr)-1] == "-" || chartArr[len(chartArr)-1] == "+" {
					chartArr = append(chartArr, v)
				} else {
					for len(chartArr) > 0 && (chartArr[len(chartArr)-1] == "*" || chartArr[len(chartArr)-1] == "/") {
						var tmp int64
						if chartArr[len(chartArr)-1] == "*" {
							tmp = numArr[len(numArr)-2] * numArr[len(numArr)-1]
						} else {
							tmp = numArr[len(numArr)-2] / numArr[len(numArr)-1]
						}

						chartArr = chartArr[0 : len(chartArr)-1]
						numArr = numArr[0 : len(numArr)-2]
						numArr = append(numArr, tmp)

					}
					chartArr = append(chartArr, v)
				}
			}
		} else {
			num, e := strconv.ParseInt(v, 10, 64)
			if e != nil {
				panic(e)
			}
			numArr = append(numArr, num)
		}
	}
	fmt.Println(numArr, chartArr)
	for len(chartArr) > 0 {
		chart := chartArr[len(chartArr)-1]

		switch chart {
		case "-", "+":
			var tmp int64
			if chartArr[len(chartArr)-1] == "-" {
				tmp = numArr[len(numArr)-2] - numArr[len(numArr)-1]
			} else {
				tmp = numArr[len(numArr)-2] + numArr[len(numArr)-1]
			}
			numArr = numArr[0 : len(numArr)-2]
			numArr = append(numArr, tmp)
		case "/", "*":
			var tmp int64
			if chartArr[len(chartArr)-1] == "*" {
				tmp = numArr[len(numArr)-2] * numArr[len(numArr)-1]
			} else {
				tmp = numArr[len(numArr)-2] / numArr[len(numArr)-1]
			}
			numArr = numArr[0 : len(numArr)-2]
			numArr = append(numArr, tmp)
		}
		chartArr = chartArr[:len(chartArr)-1]
	}
	fmt.Println(numArr)
}
