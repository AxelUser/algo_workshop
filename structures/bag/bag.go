package bag

// Bag holds collection of elements without ability to remove them.
type Bag struct {
	items []int
	size  int
}

// NewBag creates empty bag.
func NewBag() *Bag {
	return &Bag{
		items: make([]int, 0),
	}
}

// Add value to bag.
func (b *Bag) Add(item int) {
	b.items = append(b.items, item)
	b.size++
}

// IsEmpty returns true if bag is empty.
func (b *Bag) IsEmpty() bool {
	return b.size == 0
}

// Size returns number of items in bag.
func (b *Bag) Size() int {
	return b.size
}

// Iter iterates over items in bag.
func (b *Bag) Iter() <-chan int {
	ch := make(chan int)
	go func() {
		for _, item := range b.items {
			ch <- item
		}
		close(ch)
	}()

	return ch
}
