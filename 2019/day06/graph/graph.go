package graph

func New(vertices int) Graph {
	adj := make([]map[int]struct{}, 0, vertices)
	for i := 0; i < vertices; i++ {
		adj = append(adj, make(map[int]struct{}))
	}

	return Graph{
		vertices: vertices,
		adj:      adj,
	}
}

type Graph struct {
	vertices int
	adj      []map[int]struct{}
}

func (g Graph) AddEdge(v, w int) {
	g.adj[v][w] = struct{}{}
}

func (g Graph) V() int {
	return g.vertices
}

func (g Graph) Adj(v int) map[int]struct{} {
	return g.adj[v]
}
