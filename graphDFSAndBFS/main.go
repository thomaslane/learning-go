package main

import (
	"container/list"
	"fmt"

	"golang.org/x/exp/slices"
)

type Graph struct {
	Verticies map[int][]int
}

func (g *Graph) AddVertex(n int) bool {
	if _, ok := g.Verticies[n]; ok {
		fmt.Printf("Node %v already exists\n", n)
		return false
	}
	g.Verticies[n] = make([]int, 0)
	return true
}

func (g *Graph) AddEdge(vertex, node int) bool {
	if _, ok := g.Verticies[vertex]; !ok {
		fmt.Printf("Vertex %v does not exist", vertex)
		return false
	}
	if _, ok := g.Verticies[node]; !ok {
		fmt.Printf("Vertex %v does not exist", node)
		return false
	}
	if slices.Contains(g.Verticies[vertex], node) {
		fmt.Printf("Vertex %v already contains edge to %v", vertex, node)
		return false
	}
	g.Verticies[vertex] = append(g.Verticies[vertex], node)
	g.Verticies[node] = append(g.Verticies[node], vertex)
	return true
}

func (g *Graph) dfsUtil(n int, visited map[int]bool) {
	visited[n] = true
	fmt.Printf("%v ", n)

	for _, v := range g.Verticies[n] {
		if !visited[v] {
			g.dfsUtil(v, visited)
		}
	}
}

func (g *Graph) DFS(v int) {
	visited := make(map[int]bool)
	for k := range g.Verticies {
		visited[k] = false
	}

	for k := range g.Verticies {
		if !visited[k] {
			fmt.Println("---Connected---")
			g.dfsUtil(k, visited)
			fmt.Printf("\n---------------\n")
		}
	}
}

// breadth first search, takes an int which is the root node
func (g *Graph) BFS(v int) {
	if _, ok := g.Verticies[v]; !ok {
		fmt.Printf("graph does not contain node %v", v)
		return
	}

	visited := make(map[int]bool)

	queue := list.New()
	queue.PushBack(v)

	var level int

	for queue.Len() > 0 {
		// fmt.Printf("Level: %v\n", level)
		n := queue.Front()
		val := n.Value.(int)
		fmt.Printf("%v ", val)
		visited[n.Value.(int)] = true

		for _, a := range g.Verticies[val] {
			if !visited[a] {
				queue.PushBack(a)
			}
		}
		level++
		queue.Remove(n)
	}
}

func (g *Graph) BFS2(v, e int) {
	if _, ok := g.Verticies[v]; !ok {
		fmt.Printf("graph does not contain node %v", v)
		return
	}
	if _, ok := g.Verticies[e]; !ok {
		fmt.Printf("graph does not contain node %v", v)
		return
	}

	visited := make(map[int]bool)

	queue := list.New()
	queue.PushBack(v)

	pathIndex := 0
	pathList := [][]int{}

	pathList[pathIndex] = []int{v}

	for pathIndex < len(pathList) {
		// fmt.Printf("Level: %v\n", level)
		n := queue.Front()
		val := n.Value.(int)
		fmt.Printf("%v ", val)
		visited[n.Value.(int)] = true

		for _, a := range g.Verticies[val] {
			if !visited[a] {
				queue.PushBack(a)
			}
		}
		queue.Remove(n)
		pathIndex++
	}
}

func main() {
	g := &Graph{Verticies: make(map[int][]int)}

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)

	g.AddVertex(10)
	g.AddVertex(11)

	g.AddEdge(10, 11)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(4, 8)

	//fmt.Println(g)

	// g.DFS(1)
	g.BFS(7)
}
