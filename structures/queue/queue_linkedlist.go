package queue

type node struct {
	item interface{}
	next *node
}

type queueLinkedList struct {
	first  *node
	last   *node
	length int
}

// NewQueueLinkedList creates new queue based on linked-list.
func NewQueueLinkedList() Queue {
	return &queueLinkedList{}
}

func (q *queueLinkedList) Enqueue(item interface{}) {
	oldLast := q.last

	q.last = &node{
		item: item,
	}

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}

	q.length++
}

func (q *queueLinkedList) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.first.item
	q.length--
	q.first = q.first.next

	return item
}

func (q *queueLinkedList) Peek() interface{} {
	return q.first
}

func (q *queueLinkedList) IsEmpty() bool {
	return q.length == 0
}

func (q *queueLinkedList) Size() int {
	return q.length
}

func (q *queueLinkedList) Iter() <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		for !q.IsEmpty() {
			ch <- q.Dequeue()
		}

		close(ch)
	}()

	return ch
}
