package helper

import "fmt"

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
	for index > 0 {
		pIndx := h.GetParent(index)
		if h.Arr[pIndx] <= h.Arr[index] {
			break
		}
		h.Arr[pIndx], h.Arr[index] = h.Arr[index], h.Arr[pIndx]
		index = pIndx
	}
}

func (h *MinHeap) downUP(indx int) {
	l := h.GetLeftChild(indx)
	r := h.GetRightChild(indx)
	smallest := indx
	if l < h.Size && h.Arr[l] < h.Arr[smallest] {
		smallest = l
	}

	if r < h.Size && h.Arr[r] < h.Arr[smallest] {
		smallest = r
	}

	if smallest != indx {
		h.Arr[indx], h.Arr[smallest] = h.Arr[smallest], h.Arr[indx]
		h.downUP(smallest)
	}

}

func (h *MinHeap) Delete(indx int) {
	h.Arr[indx], h.Arr[h.Size-1] = h.Arr[h.Size-1], h.Arr[indx]
	h.Arr = h.Arr[:h.Size-1]
	h.Size--
	if h.Size > 0 {
		h.downUP(indx)
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.Arr)
}

func (h *MinHeap) Peek() int {
	return h.Arr[0]
}

func GetNewMinHeap(size int) *MinHeap {
	return &MinHeap{
		Arr:  make([]int, 0, size),
		Size: 0,
	}
}
