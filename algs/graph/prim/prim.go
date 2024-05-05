package prim

import (
	"graph"
	"math"
	"util"
)

type Edge struct {
	src    string
	dst    string
	weight float64
}

func MST(g *graph.Graph) []*Edge {
	visited := []int{0}
	adjVex := getAdjVex(g)
	adjMatrix := graphToAdjMatrix(g)
	lowCost := adjMatrix[0]
	idxVertexMap := getIdxVertexMap(g)

	var edges []*Edge
	for !allVertexSelected(adjVex) {
		lowestCostIdx := findLowestCostIdx(lowCost, visited)
		if lowestCostIdx < 0 {
			break
		}

		edges = append(edges, &Edge{
			src:    idxVertexMap[adjVex[lowestCostIdx]],
			dst:    idxVertexMap[lowestCostIdx],
			weight: lowCost[lowestCostIdx],
		})

		visited = append(visited, lowestCostIdx)
		compareCost := adjMatrix[lowestCostIdx]
		for i, cost := range lowCost {
			if util.Contains(visited, i) {
				continue
			}

			if compareCost[i] < cost {
				lowCost[i] = compareCost[i]
				adjVex[i] = lowestCostIdx
			}
		}
	}

	return edges
}

func allVertexSelected(adjVex []int) bool {
	for i := 1; i < len(adjVex); i++ {
		if adjVex[i] == 0 {
			return false
		}
	}
	return true
}

func findLowestCostIdx(lowCost []float64, visited []int) int {
	var lowestCost float64
	firstValue := true
	lowestCostIndex := -1
	for i := 1; i < len(lowCost); i++ {
		if util.Contains(visited, i) {
			continue
		}
		if firstValue {
			lowestCost = lowCost[i]
			lowestCostIndex = i
			firstValue = false
		} else if lowCost[i] < lowestCost {
			lowestCostIndex = i
			lowestCost = lowCost[i]
		}
	}
	return lowestCostIndex
}

func getIdxVertexMap(g *graph.Graph) map[int]string {
	m := make(map[int]string)
	for k, v := range g.Vertices() {
		m[k] = v
	}
	return m
}

func getAdjVex(g *graph.Graph) []int {
	adjVex := []int{0}
	for i := 1; i < len(g.Vertices()); i++ {
		adjVex = append(adjVex, 0)
	}
	return adjVex
}

func graphToAdjMatrix(g *graph.Graph) [][]float64 {
	vertices := g.Vertices()
	result := make([][]float64, len(vertices))
	for i := 0; i < len(vertices); i++ {
		for j := 0; j < len(vertices); j++ {
			ri := result[i]
			if len(ri) == 0 {
				result[i] = make([]float64, len(vertices))
				ri = result[i]
			}

			src := vertices[i]
			dst := vertices[j]

			if src == dst {
				ri[j] = 0.0
			} else {
				dstVertices, ok := g.Edges()[src]
				if ok {
					weight, ok := dstVertices[dst]
					if ok {
						ri[j] = weight
					} else {
						ri[j] = math.MaxFloat64
					}
				} else {
					ri[j] = math.MaxFloat64
				}
			}
		}
	}
	return result
}
