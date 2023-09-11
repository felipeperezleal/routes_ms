package src

import (
	"fmt"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
)

var route models.Routes

type Flight struct {
	Origin      string
	Destination string
	Duration    int
	Distance    float64
	Price       float64
}

func NewFlight(origin, destination string, duration int, distance, price float64) *Flight {
	return &Flight{
		Origin:      origin,
		Destination: destination,
		Duration:    duration,
		Distance:    distance,
		Price:       price,
	}
}

type Route struct {
	numNodes int
	adjList  [][]int
	flights  []*Flight
	ordering []int
}

func NewRoute(numNodes int) *Route {
	db.DBConnection()
	route = models.Routes{
		NumNodes: numNodes,
	}
	db.DB.Create(&route)
	return &Route{
		numNodes: numNodes,
		adjList:  make([][]int, numNodes),
		flights:  make([]*Flight, 0),
		ordering: make([]int, 0),
	}
}

func (g *Route) AddEdge(from, to int, flight *Flight) {
	g.adjList[from] = append(g.adjList[from], to)
	g.flights = append(g.flights, flight)
}

func (g *Route) TopoSort() []int {
	inDegree := make([]int, g.numNodes)
	for _, neighbors := range g.adjList {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	queue := []int{}
	for i, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, i)
		}
	}

	topo := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		topo = append(topo, node)

		for _, neighbor := range g.adjList[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	g.ordering = topo
	db.DB.Model(&route).Update("ordering", fmt.Sprint(topo))

	return topo
}
