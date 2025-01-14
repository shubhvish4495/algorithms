package main

import (
	"fmt"

	"github.com/shubhvish4495/algorithms/helper"
)

type MinHeap struct {
	Arr  []int
	Size int
}

func (h *MinHeap) GetLeftChild(index int) int {
	return 2*index + 1
}

func (h *MinHeap) GetRightChild(index int) int {
	return 2*index + 2
}

func (h *MinHeap) GetParent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) Insert(data int) {
	h.Arr = append(h.Arr, data)
	h.Size++
	h.Heapify(h.Size - 1)
}

func (h *MinHeap) Heapify(index int) {
	if index == 0 {
		return
	}

	pIndx := h.GetParent(index)
	if h.Arr[pIndx] > h.Arr[index] {
		h.Arr[pIndx], h.Arr[index] = h.Arr[index], h.Arr[pIndx]
		h.Heapify(pIndx)
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.Arr)
}

func main() {

	h := &MinHeap{}

	h.Insert(10)
	h.Insert(1)
	h.Insert(7)
	h.Insert(3)
	h.Insert(89)
	h.Insert(6)

	h.Print()

	h2 := helper.GetNewMinHeap(6)
	h2.Insert(10)
	h2.Insert(1)
	h2.Insert(7)
	h2.Insert(3)
	h2.Insert(89)
	h2.Insert(6)

	h2.Print()

	h2.Delete(3)
	h2.Print()

}
