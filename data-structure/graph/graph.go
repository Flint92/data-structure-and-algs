package graph

import "util"

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

func fillMap(m map[string]map[string]float64, src, dst string, weight float64) {
	dstMap, ok := m[src]
	if !ok {
		m[src] = make(map[string]float64)
		dstMap = m[src]
	}
	dstMap[dst] = weight
}
