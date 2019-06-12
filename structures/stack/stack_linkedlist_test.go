package stack

import (
	"reflect"
	"testing"
)

func TestNewStackLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want Stack
	}{
		{"Should create empty stack", &stackLinkedList{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStackLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackLinkedList_Push(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		s    *stackLinkedList
		args args
	}{
		{"Push 0", NewStackLinkedList().(*stackLinkedList), args{item: 0}},
		{"Push 100", NewStackLinkedList().(*stackLinkedList), args{item: 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.item)
			if tt.s.head.item != tt.args.item {
				t.Errorf("After stackLinkedList.Push() got %v at head, want %v", tt.s.head.item, tt.args.item)
			}
		})
	}
}

func Test_stackLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *stackLinkedList
		want interface{}
	}{
		{"Pop 10", &stackLinkedList{
			head: &node{
				item: 10,
			},
			length: 1,
		}, 10},
		{"Empty stack, should return nil", NewStackLinkedList().(*stackLinkedList), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stackLinkedList.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackLinkedList_Peek(t *testing.T) {
	tests := []struct {
		name string
		s    *stackLinkedList
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stackLinkedList.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackLinkedList_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *stackLinkedList
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("stackLinkedList.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackLinkedList_Size(t *testing.T) {
	tests := []struct {
		name string
		s    *stackLinkedList
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("stackLinkedList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackLinkedList_Iter(t *testing.T) {
	tests := []struct {
		name string
		s    *stackLinkedList
		want <-chan interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Iter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stackLinkedList.Iter() = %v, want %v", got, tt.want)
			}
		})
	}
}
