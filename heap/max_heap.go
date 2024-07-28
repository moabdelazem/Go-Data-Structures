package heap

type MaxHeap []int

// insert new data in the heap using prelocateup method
// and append the data to heap array first
func (h *MaxHeap) Insert(data int) {
	*h = append(*h, data)
	h.PrelocateUp()
}

// Given the idx i and j to swap their data in the heap
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// prelocate the node to the top of the heap
// get the last idx in the heap
// and calc the parentnode idx using `(i-1) / 2` formula
func (h *MaxHeap) PrelocateUp() {
	currentIdx := len(*h) - 1
	parentIdx := h.Parent(currentIdx)

	// if the value of the `currentnode` is bigger[maxHeap Requirment] than `parentnode` we swap them
	// and move to the parent node and repeat the procces for it parent
	for currentIdx != 0 && (*h)[currentIdx] > (*h)[parentIdx] {
		h.Swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = h.Parent(currentIdx)
	}
}

// prelocate the node to the bottom of the heap
// !Todo : Test This Function Not Sure It's Work
func (h *MaxHeap) PrelocateDown(currentIdx, endIdx int) {
	leftChildIdx := h.LeftChild(currentIdx)

	for leftChildIdx <= endIdx {
		rightChildIdx := h.RighChild(currentIdx)

		idxToSwapped := leftChildIdx

		if rightChildIdx <= endIdx && (*h)[rightChildIdx] > (*h)[idxToSwapped] {
			idxToSwapped = rightChildIdx
		}

		if (*h)[idxToSwapped] > (*h)[currentIdx] {
			h.Swap(idxToSwapped, currentIdx)
			currentIdx = idxToSwapped
			leftChildIdx = h.LeftChild(currentIdx)
		} else {
			break
		}
	}
}

// Pop The Last Node Of The Heap
func (h *MaxHeap) Delete() int {
	heapLength := len(*h) - 1
	h.Swap(0, heapLength)
	valueToRemove := (*h)[heapLength-1]
	*h = (*h)[:heapLength-1]

	h.PrelocateDown(0, heapLength-2)
	return valueToRemove
}

// Get The Length Of The Heap Array
func (h MaxHeap) Len() int {
	return len(h)
}

// Get the parent index
func (h MaxHeap) Parent(startIdx int) int {
	return (startIdx + 1) / 2
}

// Get the left child node
func (h MaxHeap) LeftChild(startIdx int) int {
	return (2 * startIdx) + 1
}

// Get the right child node
func (h MaxHeap) RighChild(startIdx int) int {
	return (2 * startIdx) + 2
}
