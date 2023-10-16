package src

type Graph struct {
	V   int
	adj map[string][]string
}

func NewGraph(v int) *Graph {
	return &Graph{
		V:   v,
		adj: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(v, w string) {
	g.adj[v] = append(g.adj[v], w)
}

func (g *Graph) TopologicalSort() []string {
	visited := make(map[string]bool)
	stack := make([]string, 0)

	var topologicalSortUtil func(string)
	topologicalSortUtil = func(v string) {
		visited[v] = true
		for _, neighbor := range g.adj[v] {
			if !visited[neighbor] {
				topologicalSortUtil(neighbor)
			}
		}
		stack = append(stack, v)
	}

	for vertex := range g.adj {
		if !visited[vertex] {
			topologicalSortUtil(vertex)
		}
	}

	reversed := make([]string, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		reversed = append(reversed, stack[i])
	}

	return reversed
}

func FindRoute(arr []string, a string, b string) []string {
	var routePath []string
	startIndex, endIndex := -1, -1

	for i, elem := range arr {
		if elem == a || elem == b {
			if startIndex == -1 {
				startIndex = i
			} else {
				endIndex = i
				break
			}
		}
	}

	if startIndex == -1 || endIndex == -1 {
		return []string{}
	}
	if startIndex > endIndex {
		startIndex, endIndex = endIndex, startIndex
	}

	for i := startIndex; i <= endIndex; i++ {
		routePath = append(routePath, arr[i])
	}

	return routePath
}
