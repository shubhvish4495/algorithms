package main

import (
	"fmt"

	"github.com/shubhvish4495/algorithms/helper"
)

type Heap struct {
	Arr  []int
	Size int
}

func (h *Heap) GetLeftChild(index int) int {
	return 2*index + 1
}

func (h *Heap) GetRightChild(index int) int {
	return 2*index + 2
}

func (h *Heap) GetParent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) Insert(value int) {
	h.Arr = append(h.Arr, value)
	h.Size++
	h.Heapify(len(h.Arr) - 1)
}

func (h *Heap) Heapify(index int) {
	if index == 0 {
		return
	}

	parIndx := h.GetParent(index)
	if h.Arr[index] > h.Arr[parIndx] {
		h.Arr[index], h.Arr[parIndx] = h.Arr[parIndx], h.Arr[index]
	}
	h.Heapify(parIndx)
}

func (h *Heap) Print() {
	fmt.Println(h.Arr)
}

func (h *Heap) downHeap(index int) {
	leftChildIndx := h.GetLeftChild(index)
	rightChildIndx := h.GetRightChild(index)

	if leftChildIndx < len(h.Arr) && h.Arr[leftChildIndx] > h.Arr[index] {
		h.Arr[leftChildIndx], h.Arr[index] = h.Arr[index], h.Arr[leftChildIndx]
		h.downHeap(leftChildIndx)
	}

	if rightChildIndx < len(h.Arr) && h.Arr[rightChildIndx] > h.Arr[index] {
		h.Arr[rightChildIndx], h.Arr[index] = h.Arr[index], h.Arr[rightChildIndx]
		h.downHeap(rightChildIndx)
	}

}

func (h *Heap) Pop() {
	if len(h.Arr) == 0 {
		fmt.Print("no element to remove")
		return
	}
	fmt.Printf("removing element from top %v", h.Arr[0])
	if len(h.Arr) == 1 {
		h.Arr = []int{}
	}
	h.Arr[0] = h.Arr[len(h.Arr)-1]
	h.Arr = h.Arr[:len(h.Arr)-1]
	h.Size--
	h.downHeap(0)

}

func (h *Heap) Delete(index int) {
	fmt.Printf("removing element %v\n", h.Arr[index])
	if len(h.Arr) == 0 && index == 0 {
		h.Arr = []int{}
	}

	if len(h.Arr)-1 == index {
		h.Arr = h.Arr[:h.Size-1]
		h.Size--
		return
	}

	h.Arr[index], h.Arr[h.Size-1] = h.Arr[h.Size-1], h.Arr[index]
	h.Arr = h.Arr[:h.Size-1]
	h.Size--

	leftChildIndx := h.GetLeftChild(index)
	rightChildIndx := h.GetRightChild(index)
	parentIndx := h.GetParent(index)

	if h.Arr[parentIndx] < h.Arr[index] {
		h.Arr[index], h.Arr[parentIndx] = h.Arr[parentIndx], h.Arr[index]
		h.Heapify(parentIndx)
	}

	if (leftChildIndx < h.Size) && !(h.Arr[leftChildIndx] < h.Arr[index]) {
		h.Arr[leftChildIndx], h.Arr[index] = h.Arr[index], h.Arr[leftChildIndx]
		h.downHeap(leftChildIndx)
	}

	if (rightChildIndx < h.Size) && !(h.Arr[rightChildIndx] < h.Arr[index]) {
		h.Arr[rightChildIndx], h.Arr[index] = h.Arr[index], h.Arr[rightChildIndx]
		h.downHeap(rightChildIndx)
	}

}

func main() {
	h := &Heap{}

	h.Insert(9)
	h.Insert(10)
	h.Insert(7)
	h.Insert(4)
	h.Insert(3)
	h.Insert(5)
	h.Insert(6)

	h.Print()

	h2 := helper.GetNewHeap()
	h2.Insert(9)
	h2.Insert(10)
	h2.Insert(7)
	h2.Insert(4)
	h2.Insert(3)
	h2.Insert(5)
	h2.Insert(6)
	h.Print()

	fmt.Println("**** Now delete ****")
	h.Delete(0)
	h.Print()

	h2.Delete(0)
	h2.Print()

}
