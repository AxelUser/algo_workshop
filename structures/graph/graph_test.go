package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 2)

	totalAdj := 0
	for range g.Adjacencies(1) {
		totalAdj++
	}

	assert.Equal(t, 3, g.CountVertices(), "Should have 3 vertices")
	assert.Equal(t, 2, totalAdj, "#1 should have 2 adjacent vertices")
}
