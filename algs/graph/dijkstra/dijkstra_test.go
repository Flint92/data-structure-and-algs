package dijkstra

import (
	"github.com/stretchr/testify/require"
	"graph"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := graph.NewGraph()

	g.AddEdgeDirected("A", "B", 4)
	g.AddEdgeDirected("A", "C", 8)
	g.AddEdgeDirected("B", "C", 11)
	g.AddEdgeDirected("B", "D", 8)
	g.AddEdgeDirected("C", "E", 7)
	g.AddEdgeDirected("D", "E", 2)
	g.AddEdgeDirected("C", "F", 1)
	g.AddEdgeDirected("E", "F", 6)
	g.AddEdgeDirected("D", "G", 4)
	g.AddEdgeDirected("F", "G", 2)
	g.AddEdgeDirected("D", "H", 7)
	g.AddEdgeDirected("H", "G", 11)
	g.AddEdgeDirected("H", "I", 9)
	g.AddEdgeDirected("G", "I", 10)

	path, dist := Dijkstra(g, "A")

	require.Equal(t, path["A"], "A")
	require.Equal(t, path["B"], "A->B")
	require.Equal(t, path["C"], "A->C")
	require.Equal(t, path["D"], "A->B->D")
	require.Equal(t, path["E"], "A->B->D->E")
	require.Equal(t, path["F"], "A->C->F")
	require.Equal(t, path["G"], "A->C->F->G")
	require.Equal(t, path["H"], "A->B->D->H")
	require.Equal(t, path["I"], "A->C->F->G->I")

	require.Equal(t, dist["A"], 0.0)
	require.Equal(t, dist["B"], 4.0)
	require.Equal(t, dist["C"], 8.0)
	require.Equal(t, dist["D"], 12.0)
	require.Equal(t, dist["E"], 14.0)
	require.Equal(t, dist["F"], 9.0)
	require.Equal(t, dist["G"], 11.0)
	require.Equal(t, dist["H"], 19.0)
	require.Equal(t, dist["I"], 21.0)
}
