package heap

type node struct {
	value int
	left  *node
	right *node
}

type minHeapLinkedList struct {
	head *node
}

func (h *minHeapLinkedList) heapify() {

}
