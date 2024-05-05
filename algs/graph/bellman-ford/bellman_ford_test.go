package bellman_ford

import (
	"github.com/stretchr/testify/require"
	"graph"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	g := graph.NewGraph()

	g.AddEdgeDirected("A", "B", 90)
	g.AddEdgeDirected("A", "C", 75)
	g.AddEdgeDirected("A", "E", 80)
	g.AddEdgeDirected("B", "D", -30)
	g.AddEdgeDirected("D", "C", 10)
	g.AddEdgeDirected("E", "D", -30)
	g.AddEdgeDirected("E", "C", 10)

	path, dist := BellmanFord(g, "A")

	require.Equal(t, path["A"], "A")
	require.Equal(t, path["B"], "A->B")
	require.Equal(t, path["C"], "A->E->D->C")
	require.Equal(t, path["D"], "A->E->D")
	require.Equal(t, path["E"], "A->E")

	require.Equal(t, dist["A"], 0.0)
	require.Equal(t, dist["B"], 90.0)
	require.Equal(t, dist["C"], 60.0)
	require.Equal(t, dist["D"], 50.0)
	require.Equal(t, dist["E"], 80.0)
}
