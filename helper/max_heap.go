package helper

import "fmt"

type MaxHeap struct {
	data   []int
	length int
}

func (h *MaxHeap) GetLeftChild(root int) int {
	return 2*root + 1
}

func (h *MaxHeap) GetRightChild(root int) int {
	return 2*root + 2
}

func (h *MaxHeap) GetParent(indx int) int {
	return (indx - 1) / 2
}

func (h *MaxHeap) Heapify(root int) {
	if root == 0 {
		return
	}

	parent := h.GetParent(root)
	if parent >= 0 && h.data[root] > h.data[parent] {
		h.data[root], h.data[parent] = h.data[parent], h.data[root]
		h.Heapify(parent)
	}
}

func (h *MaxHeap) Insert(data int) {
	h.data = append(h.data, data)
	h.length++
	h.Heapify(h.length - 1)
}

func (h *MaxHeap) Print() {
	fmt.Println(h.data)
}

func (h *MaxHeap) downUP(indx int) {
	l := h.GetLeftChild(indx)
	r := h.GetRightChild(indx)
	if l < h.length && h.data[l] > h.data[indx] {
		h.data[l], h.data[indx] = h.data[indx], h.data[l]
		h.downUP(l)
	}

	if r < h.length && h.data[r] > h.data[indx] {
		h.data[r], h.data[indx] = h.data[indx], h.data[r]
		h.downUP(r)
	}
}

func (h *MaxHeap) Delete(indx int) {
	if h.length == 0 {
		fmt.Println("delete called on empty heap")
		return
	}
	if indx == h.length-1 {
		h.data = h.data[:h.length-1]
		h.length--
		return
	}

	// if the last index remove it directly

	// if any other we replace it with last element and call heapify as required
	h.data[indx], h.data[h.length-1] = h.data[h.length-1], h.data[indx]
	h.data = h.data[:h.length-1]
	h.length--
	p := h.GetParent(indx)
	l := h.GetLeftChild(indx)
	r := h.GetRightChild(indx)

	if p >= 0 && h.data[p] > h.data[indx] {
		h.data[p], h.data[indx] = h.data[indx], h.data[p]
		h.Heapify(p)
	}

	if l < h.length && (h.data[indx] < h.data[l]) {
		h.data[indx], h.data[l] = h.data[l], h.data[indx]
		h.downUP(l)
	}
	if r < h.length && (h.data[indx] < h.data[r]) {
		h.data[indx], h.data[r] = h.data[r], h.data[indx]
		h.downUP(r)

	}

}

func (h *MaxHeap) Peak() int {
	return h.data[0]
}

func GetNewHeap() *MaxHeap {
	return &MaxHeap{
		data:   make([]int, 0),
		length: 0,
	}
}
