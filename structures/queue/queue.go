package queue

// Queue stores items at FIFO order.
type Queue interface {
	Enqueue(item interface{})
	Dequeue() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
	Iter() <-chan interface{}
}
