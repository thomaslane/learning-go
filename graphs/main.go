package main

import "fmt"

// Graph structure (adjacency list)
type Graph struct {
	verticies []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex
func (g *Graph) AddVertex(k int) {
	if contains(g.verticies, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.verticies = append(g.verticies, &Vertex{key: k})
	}
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Already exists (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.verticies {
		if v.key == k {
			return g.verticies[i]
		}
	}
	return nil
}

// Contains

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print Verticies

func (g *Graph) PrintVerticies() {
	for _, v := range g.verticies {
		fmt.Printf("\nVertex %v:", v.key)
		for _, v = range v.adjacent {
			fmt.Printf("%v ", v)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)
	test.AddEdge(3, 2)

	fmt.Println(test)
	test.PrintVerticies()
}
