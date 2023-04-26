package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {

	// 长子串
	var maxLen int = firstLen
	// 短子串
	var minLen int = secondLen
	if firstLen < secondLen {
		maxLen = secondLen
		minLen = firstLen
	}
	// 长子串index
	var maxIndex, i int
	for ; i <= len(nums)-maxLen; i++ {
		if getSum(nums, i, maxLen) > getSum(nums, maxIndex, maxLen) {
			maxIndex = i
		}
	}

	var j int
	if maxIndex+maxLen > j+minLen && j+minLen > maxIndex {
		j = maxIndex + maxLen
	}

	//	短子串index
	var minIndex int = j
	for j <= len(nums)-minLen {

		if maxIndex+maxLen > j+minLen && j+minLen > maxIndex {
			j = maxIndex + maxLen
			j++
			continue
		}
		if getSum(nums, j, minLen) > getSum(nums, minIndex, minLen) {
			minIndex = j
		}
		j++
	}
	fmt.Printf("max: %v,min: %v\r\n", maxIndex, minIndex)
	return getSum(nums, maxIndex, maxLen) + getSum(nums, minIndex, minLen)
}

func getSum(num []int, index int, len int) int {
	var sum int
	for len > 0 {
		sum += num[index]
		index++
		len--
	}
	return sum
}

func main() {
	arr := []int{8, 20, 6, 2, 20, 17, 6, 3, 20, 8, 12}

	fmt.Println(maxSumTwoNoOverlap(arr, 5, 4))
}
