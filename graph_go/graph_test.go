package main

import "testing"

func TestFindConnectedComponents(t *testing.T) {
	graph := NewGraph(8)

	graph.Connect("1", "2", 1)
	graph.Connect("2", "3", 1)
	graph.Connect("2", "4", 1)
	graph.Connect("3", "1", 1)
	graph.Connect("3", "4", 1)

	graph.Connect("4", "5", 1)
	graph.Connect("5", "4", 1)

	graph.Connect("6", "7", 1)
	graph.Connect("6", "5", 1)
	graph.Connect("7", "6", 1)
	graph.Connect("7", "8", 1)
	graph.Connect("8", "7", 1)
	graph.Connect("8", "5", 1)

	components := graph.FindConnectedComponents()

	if len(components) != 3 {
		t.Fatalf("Expected 3 components, found %d", len(components))
	}
}
