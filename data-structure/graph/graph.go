package graph

import (
	"util"
)

type Operation func(string string)

type Graph struct {
	vertices []string
	edges    map[string]map[string]float64
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[string]map[string]float64)}
}

func (g *Graph) AddEdgeDirected(src, dst string, weight float64) {
	g.addEdge(src, dst, weight, true)
}

func (g *Graph) AddEdgeUndirected(src, dst string, weight float64) {
	g.addEdge(src, dst, weight, false)
}

func (g *Graph) Vertices() []string {
	return g.vertices
}

func (g *Graph) Edges() map[string]map[string]float64 {
	return g.edges
}

func (g *Graph) DFS(op Operation) {
	visited := make(map[string]bool)
	for _, vertex := range g.Vertices() {
		dfs(vertex, g.edges, op, visited)
	}
}

func (g *Graph) DFSFrom(start string, op Operation) {
	visited := make(map[string]bool)
	dfs(start, g.edges, op, visited)
}

func (g *Graph) BFS(op Operation) {
	visited := make(map[string]bool)
	for _, vertex := range g.Vertices() {
		bfs(vertex, g.edges, op, visited)
	}
}

func (g *Graph) BFSFrom(start string, op Operation) {
	visited := make(map[string]bool)
	bfs(start, g.edges, op, visited)
}

func (g *Graph) addEdge(src, dst string, weight float64, directed bool) {
	fillMap(g.edges, src, dst, weight)
	if !directed {
		fillMap(g.edges, dst, src, weight)
	}
	g.addVertex(src)
	g.addVertex(dst)
}

func (g *Graph) addVertex(v string) {
	if !util.Contains(g.vertices, v) {
		g.vertices = append(g.vertices, v)
	}
}

func bfs(vertex string, edges map[string]map[string]float64, op Operation, visited map[string]bool) {
	queue := []string{vertex}
	for len(queue) > 0 {
		vertex = queue[0]
		queue = queue[1:]
		if visited[vertex] {
			continue
		}

		visited[vertex] = true
		op(vertex)
		dstVertices, ok := edges[vertex]
		if ok {
			for dstVertex := range dstVertices {
				queue = append(queue, dstVertex)
			}
		}
	}
}

func dfs(vertex string, edges map[string]map[string]float64, op Operation, visited map[string]bool) {
	if visited[vertex] {
		return
	}
	visited[vertex] = true
	op(vertex)
	dstVertices, ok := edges[vertex]
	if ok {
		for dstVertex := range dstVertices {
			dfs(dstVertex, edges, op, visited)
		}
	}
}

func fillMap(m map[string]map[string]float64, src, dst string, weight float64) {
	dstMap, ok := m[src]
	if !ok {
		m[src] = make(map[string]float64)
		dstMap = m[src]
	}
	dstMap[dst] = weight
}
