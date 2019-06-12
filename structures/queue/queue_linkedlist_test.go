package queue

import (
	"testing"
)

func Test_queueLinkedList_Enqueue_should_add_to_last(t *testing.T) {
	q := NewQueueLinkedList().(*queueLinkedList)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.last.item != 30 {
		t.Error("Last item should be latest inserted")
	}

	if q.first.item != 10 {
		t.Error("First item should be first inserted")
	}

	if q.Size() != 3 {
		t.Error("Should has 3 items")
	}
}

func Test_queueLinkedList_Dequeue(t *testing.T) {
	q := NewQueueLinkedList().(*queueLinkedList)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	returnedItem := q.Dequeue()

	if returnedItem != 10 {
		t.Error("Dequeued item should be first in queue")
	}

	if q.first.item != 20 {
		t.Error("First item should be next after dequeued one")
	}

	if q.Size() != 2 {
		t.Error("Should has 2 items")
	}
}

func Test_queueLinkedList_Iter_should_return_in_fifo_order(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	q := NewQueueLinkedList()

	for _, item := range items {
		q.Enqueue(item)
	}

	idx := 0
	for item := range q.Iter() {
		if items[idx] != item {
			t.Errorf("Expected %v, got %v", items[idx], item)
		}
		idx++
	}

}
