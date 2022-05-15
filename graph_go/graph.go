package main

// var emptyStruct struct{}

type Graph struct {
	edges         map[interface{}]map[interface{}]int
	inNeighbors   map[interface{}]map[interface{}]bool
	expectedNodes int
}

func NewGraph(expectedNodes int) *Graph {
	return &Graph{
		edges:         make(map[interface{}]map[interface{}]int, expectedNodes),
		inNeighbors:   make(map[interface{}]map[interface{}]bool, expectedNodes),
		expectedNodes: expectedNodes,
	}
}

func (graph *Graph) Connect(vertexFrom interface{}, vertexTo interface{}, weight int) {
	if _, ok := graph.edges[vertexFrom]; !ok {
		graph.edges[vertexFrom] = make(map[interface{}]int, graph.expectedNodes)
		graph.inNeighbors[vertexFrom] = make(map[interface{}]bool, graph.expectedNodes)
	}

	if _, ok := graph.edges[vertexTo]; !ok {
		graph.edges[vertexTo] = make(map[interface{}]int, graph.expectedNodes)
		graph.inNeighbors[vertexTo] = make(map[interface{}]bool, graph.expectedNodes)
	}

	graph.edges[vertexFrom][vertexTo] = weight
	graph.inNeighbors[vertexTo][vertexFrom] = true
}

func (graph *Graph) FindAllOutNeighbors(vertex interface{}) map[interface{}]bool {
	keys := make(map[interface{}]bool, len(graph.edges))

	for key := range graph.edges[vertex] {
		keys[key] = true
	}

	return keys
}

func (graph *Graph) FindAllInNeighbors(vertex interface{}) map[interface{}]bool {
	keys := make(map[interface{}]bool, len(graph.edges))

	for key := range graph.inNeighbors[vertex] {
		keys[key] = true
	}

	return keys
}

func (graph *Graph) FindAllVertices() map[interface{}]bool {
	keys := make(map[interface{}]bool, len(graph.edges))

	for key := range graph.edges {
		keys[key] = true
	}

	return keys
}

func (graph *Graph) visit(vertex interface{}, unvisited map[interface{}]bool, L *[]interface{}) {
	if _, ok := unvisited[vertex]; ok {
		delete(unvisited, vertex)

		for neighbor := range graph.FindAllOutNeighbors(vertex) {
			graph.visit(neighbor, unvisited, L)
		}

		*L = append(*L, vertex)
	}
}

func (graph *Graph) assign(vertex interface{}, root interface{}, components map[interface{}]map[interface{}]bool) {
	hasBeenAssigned := false

	for root := range components {
		if _, ok := components[root][vertex]; ok {
			hasBeenAssigned = true
			break
		}
	}

	if !hasBeenAssigned {
		if _, ok := components[root]; !ok {
			components[root] = make(map[interface{}]bool)
		}

		components[root][vertex] = true

		for neighbor := range graph.FindAllInNeighbors(vertex) {
			graph.assign(neighbor, root, components)
		}
	}
}

func (graph *Graph) FindConnectedComponents() map[interface{}]map[interface{}]bool {
	unvisited := graph.FindAllVertices()

	L := make([]interface{}, 0, len(unvisited))

	for vertex := range graph.FindAllVertices() {
		graph.visit(vertex, unvisited, &L)
	}

	components := make(map[interface{}]map[interface{}]bool)

	for i := len(L) - 1; i >= 0; i -= 1 {
		graph.assign(L[i], L[i], components)
	}

	return components
}
