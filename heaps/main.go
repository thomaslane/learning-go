package main

import "fmt"

// MaxHeap struct
type MaxHeap struct {
	array []int
}

// Insert adds and element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	if len(h.array) == 0 {
		err := fmt.Errorf("Can not extract from empty array")
		fmt.Println(err.Error())
		return -1
	}
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp heapifies from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown heapifies from the top to the bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1

	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of the current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get left child index
func leftChild(i int) int {
	return i*2 + 1
}

// get right child index
func rightChild(i int) int {
	return i*2 + 2
}

// swaps keys in array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 70, 30, 50, 100}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		n := m.Extract()
		fmt.Println(n, m)
	}
}
