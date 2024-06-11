package heap

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

// heap 小顶堆
type heap struct {
	data []int
	cap  uint
}

// NewHeap 初始化一个堆，如果有数据则会将传入的数据堆化
func NewHeap(v uint, data []int) *heap {
	h := new(heap)
	h.cap = v
	h.data = make([]int, 0, v)
	for i := 0; i < len(data); i++ {
		h.Add(data[i])
	}

	return h
}

// Add 向堆内插入一个元素
func (h *heap) Add(value int) bool {
	length := len(h.data)
	if h.cap == uint(length) {
		return false
	}

	// 加入到队尾然后从下往上堆化
	h.data = append(h.data, value)
	index := length
	for index > 0 {
		p := index / 2
		if h.data[index] < h.data[p] {
			h.data[index], h.data[p] = h.data[p], h.data[index]
		} else {
			break
		}
		index = p
	}

	return false
}

// Remove 移出堆顶元素，并返回
func (h *heap) Remove() (int, error) {
	length := len(h.data)
	if length == 0 {
		return 0, errors.New("heap is empty")
	}

	value := h.data[0]
	h.data[0] = h.data[length-1]
	h.data = h.data[:length-1]
	//  删除堆顶元素，将末尾的元素移到堆顶，然后从上往下堆化
	index := 0
	for index < len(h.data)/2 {
		var left, right int
		if index == 0 {
			left = 1
			right = 2
		} else {
			left = index*2 + 1
			right = index*2 + 2
		}

		minIndex := index
		if left < len(h.data) && h.data[index] > h.data[left] {
			minIndex = left
		}

		if right < len(h.data) && h.data[index] > h.data[right] {
			minIndex = right
		}

		h.data[minIndex], h.data[index] = h.data[index], h.data[minIndex]
		if minIndex == index {
			break
		}

		index = minIndex
	}

	return value, nil
}

func (h *heap) Print() {
	log.Info(h.data)
}
