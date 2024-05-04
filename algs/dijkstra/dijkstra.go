package dijkstra

import (
	"graph"
	"math"
	"sort"
	"strings"
)

// Dijkstra 求给定起始点到图中顶点最短路径
// 参数:
//
//	g (*graph.Graph): 没有负权的加权图
//	start (string): 起始点
//
// 返回值:
//
//	map[string]string: 起始点到各顶点的最短路径
//	map[string]float64: 起始点到各顶点的最短距离
func Dijkstra(g *graph.Graph, start string) (map[string]string, map[string]float64) {
	dstDist := make(map[string]float64)
	visited := make(map[string]bool)
	preNodes := make(map[string]string)

	for _, vertex := range g.Vertices() {
		dstDist[vertex] = math.MaxFloat64
	}
	dstDist[start] = 0.0

	startNode := &Item{item: start, priority: 0.0}

	queue := []*Item{startNode}

	for len(queue) > 0 {
		sortByPriority(queue)

		node := queue[0]
		queue = queue[1:]

		if visited[node.item] {
			continue
		}

		visited[node.item] = true

		for neighbor, weight := range g.Edges()[node.item] {
			if dstDist[node.item]+weight < dstDist[neighbor] {
				dstDist[neighbor] = dstDist[node.item] + weight
				preNodes[neighbor] = node.item
				queue = append(queue, &Item{item: neighbor, priority: dstDist[neighbor]})
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

type Item struct {
	item     string
	priority float64
}

func sortByPriority(items []*Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].priority > items[j].priority
	})
}
