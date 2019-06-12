package stack

type stackLinkedList struct {
	head   *node
	length int
}

type node struct {
	item interface{}
	next *node
}

// NewStackLinkedList creates new stack which is based on linked-list.
func NewStackLinkedList() Stack {
	return &stackLinkedList{}
}

func (s *stackLinkedList) Push(item interface{}) {
	newHead := &node{
		item: item,
		next: s.head,
	}

	s.head = newHead
	s.length++
}

func (s *stackLinkedList) Pop() interface{} {
	if s.IsEmpty() {
		return nil // Maybe return error
	}

	item := s.head.item
	s.head = s.head.next
	s.length--

	return item
}

func (s *stackLinkedList) Peek() interface{} {
	if s.IsEmpty() {
		return nil // Maybe return error
	}

	return s.head.item
}

func (s *stackLinkedList) IsEmpty() bool {
	return s.length == 0
}

func (s *stackLinkedList) Size() int {
	return s.length
}

func (s *stackLinkedList) Iter() <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		for !s.IsEmpty() {
			ch <- s.Pop()
		}

		close(ch)
	}()

	return ch
}
