package main

var emptyStruct struct{}

type Graph struct {
	edges       map[interface{}]map[interface{}]int
	inNeighbors map[interface{}]map[interface{}]struct{}
}

func NewGraph(expectedNodes int) *Graph {
	return &Graph{
		edges:       make(map[interface{}]map[interface{}]int, expectedNodes),
		inNeighbors: make(map[interface{}]map[interface{}]struct{}, expectedNodes),
	}
}

func (graph *Graph) Connect(vertexFrom interface{}, vertexTo interface{}, weight int) {
	if _, ok := graph.edges[vertexFrom]; !ok {
		graph.edges[vertexFrom] = make(map[interface{}]int, len(graph.edges))
		graph.inNeighbors[vertexFrom] = make(map[interface{}]struct{}, len(graph.inNeighbors))
	}

	if _, ok := graph.edges[vertexTo]; !ok {
		graph.edges[vertexTo] = make(map[interface{}]int, len(graph.edges))
		graph.inNeighbors[vertexTo] = make(map[interface{}]struct{}, len(graph.inNeighbors))
	}

	graph.edges[vertexFrom][vertexTo] = weight
	graph.inNeighbors[vertexTo][vertexFrom] = emptyStruct
}

func (graph *Graph) FindAllOutNeighbors(vertex interface{}) map[interface{}]struct{} {
	keys := make(map[interface{}]struct{}, len(graph.edges))

	for key := range graph.edges[vertex] {
		keys[key] = emptyStruct
	}

	return keys
}

func (graph *Graph) FindAllInNeighbors(vertex interface{}) map[interface{}]struct{} {
	keys := make(map[interface{}]struct{}, len(graph.edges))

	for key := range graph.inNeighbors[vertex] {
		keys[key] = emptyStruct
	}

	return keys
}

func (graph *Graph) FindAllVertices() map[interface{}]struct{} {
	keys := make(map[interface{}]struct{}, len(graph.edges))

	for key := range graph.edges {
		keys[key] = emptyStruct
	}

	return keys
}

func (graph *Graph) FindConnectedComponents() map[interface{}]map[interface{}]struct{} {
	var visit func(interface{}, map[interface{}]struct{}, *[]interface{})
	var assign func(interface{}, interface{}, map[interface{}]map[interface{}]struct{})

	visit = func(vertex interface{}, unvisited map[interface{}]struct{}, L *[]interface{}) {
		if _, ok := unvisited[vertex]; ok {
			delete(unvisited, vertex)

			for neighbor := range graph.FindAllOutNeighbors(vertex) {
				visit(neighbor, unvisited, L)
			}

			*L = append(*L, vertex)
		}
	}

	assign = func(vertex interface{}, root interface{}, components map[interface{}]map[interface{}]struct{}) {
		hasBeenAssigned := false

		for root := range components {
			if _, ok := components[root][vertex]; ok {
				hasBeenAssigned = true
				break
			}
		}

		if !hasBeenAssigned {
			if _, ok := components[root]; !ok {
				components[root] = make(map[interface{}]struct{})
			}

			components[root][vertex] = emptyStruct

			for neighbor := range graph.FindAllInNeighbors(vertex) {
				assign(neighbor, root, components)
			}
		}
	}

	unvisited := graph.FindAllVertices()

	L := make([]interface{}, 0, len(unvisited))

	for vertex := range graph.FindAllVertices() {
		visit(vertex, unvisited, &L)
	}

	components := make(map[interface{}]map[interface{}]struct{})

	for i := len(L) - 1; i >= 0; i -= 1 {
		assign(L[i], L[i], components)
	}

	return components
}
