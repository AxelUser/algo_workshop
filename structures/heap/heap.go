package heap

// MinHeap represents heap that stores minimum value first.
type MinHeap interface {
	Insert(key int)
	ExtractMin() int
	IsEmpty() bool
	DecreaseKey(key int)
}
