package stack

// Stack stores items in LIFO order
type Stack interface {
	Push(item interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
	Iter() <-chan interface{}
}
