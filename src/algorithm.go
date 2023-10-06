package src

import (
	"fmt"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
)

func NewFlight(origin, destination string, duration int, price float64) *models.Flight {
	if duration < 0 || price < 0 {
		return nil
	}

	return &models.Flight{
		Origin:      origin,
		Destination: destination,
		Duration:    duration,
		Price:       price,
	}
}

type Route struct {
	numNodes int
	adjList  [][]int
	flights  []*models.Flight
	ordering []int
}

func NewRoute(numNodes int) (*Route, *models.Routes) {
	db.DBConnection()
	route := models.Routes{
		NumNodes: numNodes,
	}
	db.DB.Create(&route)
	return &Route{
		numNodes: numNodes,
		adjList:  make([][]int, numNodes),
		flights:  make([]*models.Flight, 0),
		ordering: make([]int, 0),
	}, &route
}

func (g *Route) AddEdge(from, to int, flight *models.Flight) {
	if from < 0 || from >= g.numNodes || to < 0 || to >= g.numNodes {
		return
	}

	g.adjList[from] = append(g.adjList[from], to)
	g.flights = append(g.flights, flight)
}

func (g *Route) TopoSort(dbRoute *models.Routes) []int {
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

	// Actualizar el orden en la base de datos.
	db.DB.Model(dbRoute).Update("ordering", fmt.Sprint(topo))

	return topo
}
