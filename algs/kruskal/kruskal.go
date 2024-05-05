package kruskal

import (
	"graph"
	"sort"
)

type Edge struct {
	src    string
	dst    string
	weight float64
}

func MST(g *graph.Graph) []*Edge {
	edgeSetArray := graphToEdgeSetArray(g)
	sets := graphToSet(g)

	sortByWeight(edgeSetArray)

	var edges []*Edge
	for len(edgeSetArray) > 0 {
		if allConnected(sets) {
			break
		}

		edgeSet := edgeSetArray[0]
		edgeSetArray = edgeSetArray[1:]

		if sets[edgeSet.src] != sets[edgeSet.dst] {
			edges = append(edges, edgeSet)
			union(edgeSet.src, edgeSet.dst, sets)
		}
	}

	return edges
}

func allConnected(sets map[string]string) bool {
	var firstValue string
	first := true

	for _, value := range sets {
		if first {
			firstValue = value
			first = false
		} else {
			if value != firstValue {
				return false
			}
		}
	}

	return true
}

func union(src, dst string, sets map[string]string) {
	var modifiedKeys []string

	dstVal := sets[dst]
	srcVal := sets[src]

	for key, value := range sets {
		if value == dstVal {
			modifiedKeys = append(modifiedKeys, key)
		}
	}

	if len(modifiedKeys) > 0 {
		for _, key := range modifiedKeys {
			sets[key] = srcVal
		}
	}
}

func graphToSet(g *graph.Graph) map[string]string {
	sets := make(map[string]string)
	for _, vertex := range g.Vertices() {
		sets[vertex] = vertex
	}
	return sets
}

func graphToEdgeSetArray(g *graph.Graph) []*Edge {
	var edgeSetArray []*Edge

	visited := make(map[string]bool)
	for _, vertex := range g.Vertices() {
		dstVertices, ok := g.Edges()[vertex]
		visited[vertex] = true
		if ok {
			for dstVertex, weight := range dstVertices {
				if !visited[dstVertex] {
					edgeSetArray = append(edgeSetArray, &Edge{src: vertex, dst: dstVertex, weight: weight})
				}
			}
		}
	}

	return edgeSetArray
}

func sortByWeight(items []*Edge) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].weight < items[j].weight
	})
}
