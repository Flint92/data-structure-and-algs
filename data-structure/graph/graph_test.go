package graph

import (
	"github.com/stretchr/testify/require"
	"testing"
	"util"
)

func TestGraph_AddEdgeDirected(t *testing.T) {
	g := NewGraph()
	g.AddEdgeDirected("A", "B", 1)
	g.AddEdgeDirected("A", "D", 1)
	g.AddEdgeDirected("D", "B", 1)
	g.AddEdgeDirected("D", "E", 1)
	g.AddEdgeDirected("B", "E", 1)
	g.AddEdgeDirected("B", "C", 1)
	g.AddEdgeDirected("E", "C", 1)

	vertices := g.Vertices()
	require.Equal(t, 5, len(vertices))
	require.True(t, util.Contains(vertices, "A"))
	require.True(t, util.Contains(vertices, "B"))
	require.True(t, util.Contains(vertices, "C"))
	require.True(t, util.Contains(vertices, "D"))
	require.True(t, util.Contains(vertices, "E"))

}

func TestGraph_DFS(t *testing.T) {
	g := NewGraph()
	g.AddEdgeDirected("A", "B", 1)
	g.AddEdgeDirected("A", "D", 1)
	g.AddEdgeDirected("D", "B", 1)
	g.AddEdgeDirected("D", "E", 1)
	g.AddEdgeDirected("B", "E", 1)
	g.AddEdgeDirected("B", "C", 1)
	g.AddEdgeDirected("E", "F", 1)

	var dfsPath []string
	g.DFSFrom("A", func(v string) {
		dfsPath = append(dfsPath, v)
	})

	require.Equal(t, []string{"A", "B", "E", "F", "C", "D"}, dfsPath)
}

func TestGraph_BFS(t *testing.T) {
	g := NewGraph()
	g.AddEdgeDirected("A", "B", 1)
	g.AddEdgeDirected("A", "D", 1)
	g.AddEdgeDirected("D", "B", 1)
	g.AddEdgeDirected("D", "E", 1)
	g.AddEdgeDirected("B", "E", 1)
	g.AddEdgeDirected("B", "C", 1)
	g.AddEdgeDirected("E", "F", 1)

	var bfsPath []string
	g.BFSFrom("A", func(v string) {
		bfsPath = append(bfsPath, v)
	})

	require.Equal(t, []string{"A", "B", "D", "E", "C", "F"}, bfsPath)
}
