package main

import (
	"fmt"

	"github.com/moabdelazem/Go-Data-Structures/heap"
)

func main() {
	// Testing The Heap
	var minHeap heap.MinHeap

	minHeap.InsertNode(9)
	minHeap.InsertNode(31)
	minHeap.InsertNode(40)
	minHeap.InsertNode(22)
	minHeap.InsertNode(10)
	minHeap.InsertNode(15)
	minHeap.InsertNode(1)
	minHeap.InsertNode(25)
	minHeap.InsertNode(91)

	fmt.Println(minHeap)

	minHeap.Delete()
	fmt.Println(minHeap)
}
