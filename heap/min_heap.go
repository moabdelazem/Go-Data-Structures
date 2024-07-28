package heap

// The Main Heap Type
type MinHeap []int

// Get the length of the heap
func (h MinHeap) Len() int {
	return len(h)
}

// Swap The `i` idx with the `j` idx
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Insert to the heap
func (h *MinHeap) Insert(data int) {
	*h = append(*h, data)
	h.PrelocateUp()
}

// Prelocate Up The Node (SiftUp)
func (h *MinHeap) PrelocateUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2

	// If The child value is smaller than the parent child we will swap them
	for currentIdx != 0 && (*h)[currentIdx] < (*h)[parentIdx] {
		h.Swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Prelocate Down The Node (SiftDown)
func (h *MinHeap) PrelocateDown(currentIdx, endIdx int) {
	leftChildIdx := (2 * currentIdx) + 1

	for leftChildIdx <= endIdx {
		rightChildIdx := (2 * currentIdx) + 2

		// Get the smaller noded to be swapped
		idxToBeSwapped := leftChildIdx
		if (*h)[rightChildIdx] < (*h)[leftChildIdx] {
			idxToBeSwapped = rightChildIdx
		}

		// Check if the index to be swapped is smaller than value of the current index
		if (*h)[idxToBeSwapped] < (*h)[currentIdx] {
			h.Swap(idxToBeSwapped, currentIdx)
			currentIdx = idxToBeSwapped
			leftChildIdx = (2 * currentIdx) + 1
		}
	}
}

// Pop The Last Node In The Min Heap
func (h *MinHeap) Delete() int {
	heapLength := len(*h) - 1
	h.Swap(0, heapLength)
	valueToRemove := (*h)[heapLength-1]
	*h = (*h)[:heapLength-1]

	h.PrelocateDown(0, heapLength-2)
	return valueToRemove
}
