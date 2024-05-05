package kruskal

import (
	"github.com/stretchr/testify/require"
	"graph"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	g := graph.NewGraph()
	g.AddEdgeUndirected("A", "B", 6)
	g.AddEdgeUndirected("A", "C", 1)
	g.AddEdgeUndirected("A", "D", 5)
	g.AddEdgeUndirected("B", "C", 5)
	g.AddEdgeUndirected("B", "E", 3)
	g.AddEdgeUndirected("C", "D", 5)
	g.AddEdgeUndirected("C", "E", 6)
	g.AddEdgeUndirected("C", "F", 4)
	g.AddEdgeUndirected("D", "F", 2)
	g.AddEdgeUndirected("E", "F", 6)

	edges := MST(g)
	require.Equal(t, 5, len(edges))
	require.Equal(t, &Edge{src: "A", dst: "C", weight: 1.0}, edges[0])
	require.Equal(t, &Edge{src: "D", dst: "F", weight: 2.0}, edges[1])
	require.Equal(t, &Edge{src: "B", dst: "E", weight: 3.0}, edges[2])
	require.Equal(t, &Edge{src: "C", dst: "F", weight: 4.0}, edges[3])
	require.Equal(t, &Edge{src: "B", dst: "C", weight: 5.0}, edges[4])
}
