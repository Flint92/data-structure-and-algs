package bellman_ford

import (
	"graph"
	"math"
	"strings"
)

func BellmanFord(g *graph.Graph, start string) (map[string]string, map[string]float64) {
	dstDist := make(map[string]float64)
	preNodes := make(map[string]string)

	for _, vertex := range g.Vertices() {
		dstDist[vertex] = math.MaxFloat64
	}
	dstDist[start] = 0.0

	for _, vertex := range g.Vertices() {
		dstVertices, ok := g.Edges()[vertex]
		if ok {
			for neighbor, weight := range dstVertices {
				if dstDist[vertex]+weight < dstDist[neighbor] {
					dstDist[neighbor] = dstDist[vertex] + weight
					preNodes[neighbor] = vertex
				}
			}
		}
	}

	dstPath := make(map[string]string)
	for dst := range dstDist {
		paths := []string{dst}

		for pre, ok := preNodes[dst]; ok; pre, ok = preNodes[pre] {
			paths = append([]string{pre}, paths...)
		}

		dstPath[dst] = strings.Join(paths, "->")
	}

	return dstPath, dstDist
}
