package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sut := NewBag()
	sut.Add(54)

	assert.Equal(t, sut.items[0], 54, "Bag should contain 54")
}
