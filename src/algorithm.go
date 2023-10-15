package src

type Graph struct {
	Edges   map[string][]string
	Sorted  []string
	Visited map[string]bool
}

func NewGraph() *Graph {
	return &Graph{
		Edges:   make(map[string][]string),
		Visited: make(map[string]bool),
	}
}

func (g *Graph) AddRoute(origin, destiny string) {
	g.Edges[origin] = append(g.Edges[origin], destiny)
}

func (g *Graph) TopologicalSort(node string) {
	if !g.Visited[node] {
		g.Visited[node] = true
		for _, neighbor := range g.Edges[node] {
			g.TopologicalSort(neighbor)
		}
		g.Sorted = append(g.Sorted, node)
	}
}
