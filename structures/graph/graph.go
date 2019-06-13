package graph

import (
	"bytes"
	"fmt"

	"github.com/AxelUser/algo_workshop/structures/bag"
)

// Edge contains weight and index of adjacent node
type Edge struct {
	ToIndex int
	Weight  int
}

func NewEdge(idx int, w int) *Edge {
	return &Edge{
		ToIndex: idx,
		Weight:  w,
	}
}

// Graph holds nodes and edges.
type Graph struct {
	nodes map[int]*bag.Bag
}

// NewGraph creates new graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*bag.Bag, 0),
	}
}

func (g *Graph) AddNode(index int) {
	if _, ok := g.nodes[index]; !ok {
		g.nodes[index] = bag.NewBag()
	}
}

func (g *Graph) AddEdge(v int, w int, weight int) {
	g.AddNode(v)
	g.nodes[v].Add(NewEdge(w, weight))

	g.AddNode(w)
	g.nodes[w].Add(NewEdge(v, weight))
}

func (g *Graph) Adjacencies(v int) <-chan int {
	ch := make(chan int)

	go func() {
		for item := range g.nodes[v].Iter() {
			ch <- item.(*Edge).ToIndex
		}

		close(ch)
	}()

	return ch
}

func (g *Graph) Edges(v int) <-chan *Edge {
	ch := make(chan *Edge)

	go func() {
		for e := range g.nodes[v].Iter() {
			ch <- e.(*Edge)
		}
		close(ch)
	}()

	return ch
}

func (g *Graph) FindMinEdge(v int) *Edge {
	var found *Edge

	for e := range g.Edges(v) {
		if found == nil || e.Weight < found.Weight {
			found = e
		}
	}

	return found
}

func (g *Graph) CountVertices() int {
	return len(g.nodes)
}

func (g *Graph) CountEdges() int {
	total := 0
	for _, v := range g.nodes {
		total += v.Size()
	}

	return total
}

func (g *Graph) String() string {
	var buffer bytes.Buffer
	for v, n := range g.nodes {
		for e := range n.Iter() {
			buffer.WriteString(fmt.Sprintf("%v - %v", v, e.(Edge).ToIndex))
		}
	}

	return buffer.String()
}
